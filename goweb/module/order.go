package module

import (
	"bytes"
	"encoding/json"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册订单路由
func RegisterOrderRoutes(router *gin.Engine) {
	router.POST("/order-list", getOrderList)
	router.POST("/supplier-order-list", getSupplierOrderList)
	router.POST("/agent-order-notice", agentOrderNotice)
	router.PUT("/update-order-remark", updateOrderRemark)
	router.POST("/backup-submit-log", backupSubmitLog)
	router.DELETE("/delete-backup-submit-log", deleteBackupSubmitLog)
	router.POST("/agent-order-submit-log", agentOrderSubmitLog)
	router.POST("/agent-order-query-log", agentOrderQueryLog)
	router.POST("/agent-order-notice-log", agentOrderNoticeLog)
	router.PUT("/batch-update-order-status-remark", batchUpdateOrderStatusRemark)
	router.POST("/batch-order_notice", batchOrderNotice)
	router.PUT("/batch-backup-submit", batchBackupSubmit)
	router.PUT("/batch-backup-cancel", batchBackupCancel)
	router.PUT("/batch-order-timeout", batchOrderTimeout)
	router.PUT("/batch-order-cancel",batchOrderCancel)
	router.PUT("/batch-order-manual",batchOrderManual)
}

// @Tags 订单
// @Summary 批量手动提交订单
func batchOrderManual(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	for _, id := range ids {
		order_id := id.(float64)
		var order model.Order_list
		model.Db.Where("id = ?", order_id).First(&order)
		if order.Status == 0 {
			order.Status = 1
			model.Db.Save(&order)
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量订单取消
func batchOrderCancel(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	for _, id := range ids {
		order_id := id.(float64)
		var order model.Order_list
		model.Db.Where("id = ?", order_id).First(&order)
		if order.Status == 1 {
			order.Status = 0
			model.Db.Save(&order)
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量订单超时
func batchOrderTimeout(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	operType := requestData["type"].(float64)
	seconds := requestData["seconds"].(float64)
	create_after := requestData["create_after"].(float64)
	timePoint := requestData["time"].(string)
	ids := requestData["ids"].([]interface{})
	// 操作类型包括四种，0：立即超时，1：多久后超时，2：指定时间超时，3：订单创建多久后超时
	switch operType {
	case 0:
		for _, id := range ids {
			order_id := id.(float64)
			var order model.Order_list
			model.Db.Where("id = ?", order_id).First(&order)
			if order.Is_timeout == 0 {
				order.Is_timeout = 1
				model.Db.Save(&order)
			}
		}
	case 1:
		// 多久后超时,单位是秒，创建一个携程执行超时操作
		go func() {
			//等待seconds秒后执行
			time.Sleep(time.Duration(seconds) * time.Second)
			for _, id := range ids {
				order_id := id.(float64)
				var order model.Order_list
				model.Db.Where("id = ?", order_id).First(&order)
				if order.Is_timeout == 0 {
					order.Is_timeout = 1
					model.Db.Save(&order)
				}
			}
		}()
	case 2:
		// 指定时间超时
		go func() {
			//等待指定时间后执行
			t, _ := time.Parse("2006-01-02 15:04:05", timePoint)
			time.Sleep(time.Until(t))
			for _, id := range ids {
				order_id := id.(float64)
				var order model.Order_list
				model.Db.Where("id = ?", order_id).First(&order)
				if order.Is_timeout == 0 {
					order.Is_timeout = 1
					model.Db.Save(&order)
				}
			}
		}()
	case 3:
		// 订单创建多久后超时
		for _, id := range ids {
			go func() {
				order_id := id.(float64)
				var order model.Order_list
				model.Db.Where("id = ?", order_id).First(&order)
				if order.Is_timeout == 0 {
					t, _ := time.Parse("2006-01-02 15:04:05", order.Create_time.Format("2006-01-02 15:04:05"))
					t = t.Add(time.Duration(create_after) * time.Second)
					time.Sleep(time.Until(t))
					order.Is_timeout = 1
					model.Db.Save(&order)
				}
			}()
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量备用通道取消
func batchBackupCancel(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	for _, id := range ids {
		order_id := id.(float64)
		var order model.Order_list
		model.Db.Where("id = ?", order_id).First(&order)
		if order.Status == 1 {
			order.Status = 0
			model.Db.Save(&order)
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量备用通道重新提交
func batchBackupSubmit(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	// ids := requestData["ids"].([]interface{})
	// count := requestData["count"].(float64)
	// interval := requestData["interval"].(float64)
	// for _, id := range ids {

	// }
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量通知订单
func batchOrderNotice(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	for _, id := range ids {
		order_id := id.(float64)
		var order model.Order_list
		model.Db.Where("id = ?", order_id).First(&order)
		var agent model.Agent_account
		model.Db.Where("id = ?", order.Agent_id).First(&agent)
		if agent.Notification_address != "" {
			postData := map[string]interface{}{
				"order_id":     order.ID,
				"order_status": order.Status,
				"agent_id":     agent.Id,
			}
			jsonBytes, _ := json.Marshal(postData)
			resp, err := http.Post(agent.Notification_address, "application/json", bytes.NewBuffer(jsonBytes))
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"success": false, "message": "notify failed"})
				return
			}
			defer resp.Body.Close()
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 批量更新订单状态和备注
func batchUpdateOrderStatusRemark(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	status := requestData["status"].(float64)
	remark := requestData["remark"].(string)
	model.Db.Model(&model.Order_list{}).Where("id in (?)", ids).Updates(map[string]interface{}{"status": status, "remark": remark})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 代理商订单通知日志
func agentOrderNoticeLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	db := model.Db
	db = db.Where("order_id = ?", order_id)
	result := []model.Order_notify_log{}
	db.Model(&model.Order_notify_log{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

// @Tags 订单
// @Summary 代理商订单查询日志
func agentOrderQueryLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	db := model.Db
	db = db.Where("order_id = ?", order_id)
	result := []model.Order_query_log{}
	db.Model(&model.Order_query_log{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

// @Tags 订单
// @Summary 代理商订单提交日志
func agentOrderSubmitLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	db := model.Db
	db = db.Where("order_id = ?", order_id)
	result := []model.Order_submit_log{}
	db.Model(&model.Order_submit_log{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

// @Tags 订单
// @Summary 删除备份提交日志
func deleteBackupSubmitLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	id := requestData["id"].(float64)
	db := model.Db
	db = db.Where("id = ?", id)
	db.Delete(&model.Order_backup_submit_log{})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 备份提交日志
func backupSubmitLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	db := model.Db
	db = db.Where("order_id = ?", order_id)
	result := []model.Order_backup_submit_log{}
	db.Model(&model.Order_backup_submit_log{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": gin.H{
		"list":        result,
		"total":       len(result),
		"currentPage": 1,
		"pageSize":    len(result),
	}})

}

// @Tags 订单
// @Summary 更新订单备注
func updateOrderRemark(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	remark := requestData["remark"].(string)
	model.Db.Model(&model.Order_list{}).Where("id = ?", order_id).Update("remark", remark)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// @Tags 订单
// @Summary 代理商订单通知
func agentOrderNotice(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	order_id := requestData["order_id"].(float64)
	agent_id := requestData["agent_id"].(float64)
	// 在订单表中以订单ID查找到订单状态，在代理商表中以代理商ID查找到代理商信息的通知地址，将订单状态通知给代理商
	var order model.Order_list
	model.Db.Where("id = ?", order_id).First(&order)
	var agent model.Agent_account
	model.Db.Where("id = ?", agent_id).First(&agent)
	// 通知代理商,通知方式是http发送post请求到代理商的通知地址
	if agent.Notification_address != "" {
		postData := map[string]interface{}{
			"order_id":     order.ID,
			"order_status": order.Status,
			"agent_id":     agent.Id,
		}
		jsonBytes, _ := json.Marshal(postData)
		resp, err := http.Post(agent.Notification_address, "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "notify failed"})
			return
		}
		defer resp.Body.Close()
		c.JSON(http.StatusOK, gin.H{"success": true, "message": resp.Status})
	}
}

// @Tags 订单
// @Summary 获取供应商订单列表
func getSupplierOrderList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	business_type := requestData["business_type"]
	switch business_type.(type) {
	case float64:
		db = db.Where("business_type = ?", business_type)
	}

	order_id := requestData["order_id"]
	if order_id != nil && order_id!= "" {
		order_id = order_id.(float64)
		db = db.Where("order_id = ?", order_id)
	}

	id := requestData["id"]
	if id != nil && id != "" {
		idStr := strings.ReplaceAll(id.(string), " ", ",")
		ls := strings.Split(idStr, ",")
		db = db.Where("id in (?)", ls)
	}

	up_id := requestData["up_id"]
	if up_id != nil && up_id != "" {
		db = db.Where("up_id like ?", "%"+up_id.(string)+"%")
	}

	recharge_number := requestData["recharge_number"]
	if recharge_number != nil && recharge_number != "" {
		recharge_numberStr := strings.ReplaceAll(recharge_number.(string), " ", ",")
		ls := strings.Split(recharge_numberStr, ",")
		db = db.Where("recharge_number in (?)", ls)
	}

	agent_id := requestData["agent_id"]
	switch agent_id.(type) {
	case float64:
		db = db.Where("agent_id = ?", agent_id)
	}

	product_category := requestData["product_category"]
	switch product_category.(type) {
	case float64:
		db = db.Where("product_category = ?", product_category)
	}

	product_id := requestData["product_id"]
	switch product_id.(type) {
	case float64:
		db = db.Where("product_id = ?", product_id)
	}

	base_price := requestData["base_price"]
	switch base_price.(type) {
	case float64:
		db = db.Where("base_price = ?", base_price)
	}

	remark := requestData["remark"]
	if remark != nil && remark != "" {
		db = db.Where("remark like ?", "%"+remark.(string)+"%")
	}

	status := requestData["status"]
	switch status.(type) {
	case float64:
		db = db.Where("status = ?", status)
	}

	supplier := requestData["supplier"]
	switch supplier.(type) {
	case float64:
		db = db.Where("supplier = ?", supplier)
	}

	create_time := requestData["create_time"]
	if create_time != nil && create_time != "" {
		create_timeList := strings.Split(create_time.(string), "-")
		if len(create_timeList) == 2 {
			db = db.Where("create_time between ? and ?", create_timeList[0], create_timeList[1])
		}
	}

	finish_time := requestData["finish_time"]
	if finish_time != nil && finish_time != "" {
		finish_timeList := strings.Split(finish_time.(string), "-")
		if len(finish_timeList) == 2 {
			db = db.Where("finish_time between ? and ?", finish_timeList[0], finish_timeList[1])
		}
	}

	order_time := requestData["order_time"]
	if order_time != nil && order_time != "" {
		order_timeList := strings.Split(order_time.(string), "-")
		if len(order_timeList) == 2 {
			db = db.Where("create_time between ? and ?", order_timeList[0], order_timeList[1])
		}
	}

	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)

	offset := (page - 1) * pageSize
	var total int64
	var result []model.Order_supplier
	db.Model(&model.Order_supplier{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": gin.H{
		"list":        result,
		"total":       total,
		"currentPage": page,
		"pageSize":    pageSize,
	}})
}

// @Tags 订单
// @Summary 获取订单列表
func getOrderList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	business_type := requestData["business_type"]
	switch business_type.(type) {
	case float64:
		db = db.Where("business_type = ?", business_type)
	}
	// 根据订单ID查询,如果有多个订单ID,则查询多个订单,ID之间用逗号分隔或者空格隔开
	id := requestData["id"]
	if id != nil && id != "" {
		idStr := strings.ReplaceAll(id.(string), " ", ",")
		ls := strings.Split(idStr, ",")
		db = db.Where("id in (?)", ls)
	}

	down_id := requestData["down_id"]
	if down_id != nil && down_id != "" {
		down_idStr := strings.ReplaceAll(down_id.(string), " ", ",")
		ls := strings.Split(down_idStr, ",")
		db = db.Where("down_id in (?)", ls)
	}

	notify_status := requestData["notify_status"]
	switch notify_status.(type) {
	case float64:
		db = db.Where("notify_status = ?", notify_status)
	}

	recharge_number := requestData["recharge_number"]
	if recharge_number != nil && recharge_number != "" {
		recharge_numberStr := strings.ReplaceAll(recharge_number.(string), " ", ",")
		ls := strings.Split(recharge_numberStr, ",")
		db = db.Where("recharge_number in (?)", ls)
	}

	agent_id := requestData["agent_id"]
	switch agent_id.(type) {
	case float64:
		db = db.Where("agent_id = ?", agent_id)
	}

	product_category := requestData["product_category"]
	switch product_category.(type) {
	case float64:
		db = db.Where("product_category = ?", product_category)
	}

	product_id := requestData["product_id"]
	switch product_id.(type) {
	case float64:
		db = db.Where("product_id = ?", product_id)
	}

	base_price := requestData["base_price"]
	switch base_price.(type) {
	case float64:
		db = db.Where("base_price = ?", base_price)
	}

	remark := requestData["remark"]
	if remark != nil && remark != "" {
		db = db.Where("remark like ?", "%"+remark.(string)+"%")
	}

	status := requestData["status"]
	switch status.(type) {
	case float64:
		db = db.Where("status = ?", status)
	}

	is_timeout := requestData["is_timeout"]
	switch is_timeout.(type) {
	case float64:
		db = db.Where("is_timeout = ?", is_timeout)
	}

	is_cancel := requestData["is_cancel"]
	switch is_cancel.(type) {
	case float64:
		db = db.Where("is_cancel = ?", is_cancel)
	}

	location := requestData["location"]
	if location != nil && location != "" {
		db = db.Where("location like ?", "%"+location.(string)+"%")
	}

	special_params := requestData["special_params"]
	if special_params != nil && special_params != "" {
		db = db.Where("special_params like ?", "%"+special_params.(string)+"%")
	}

	create_time := requestData["create_time"]
	if create_time != nil && create_time != "" {
		create_timeList := strings.Split(create_time.(string), "-")
		if len(create_timeList) == 2 {
			db = db.Where("create_time between ? and ?", create_timeList[0], create_timeList[1])
		}
	}

	finish_time := requestData["finish_time"]
	if finish_time != nil && finish_time != "" {
		finish_timeList := strings.Split(finish_time.(string), "-")
		if len(finish_timeList) == 2 {
			db = db.Where("finish_time between ? and ?", finish_timeList[0], finish_timeList[1])
		}
	}

	offset := (page - 1) * pageSize
	var total int64
	var result []model.Order_list
	db.Model(&model.Order_list{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": gin.H{
		"list":        result,
		"total":       total,
		"currentPage": page,
		"pageSize":    pageSize,
	}})
}
