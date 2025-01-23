package module

import (
	"fmt"
	"main/middleware"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 注册代理商接口路由
func RegisterAgentRoutes(router *gin.Engine) {
	router.POST("/agent-list", getAgentList)
	router.POST("/agent", addAgent)
	router.PUT("/agent", updateAgent)
	router.DELETE("/agent", deleteAgent)
	router.PUT("/batch-update-agent-status", batchUpdateAgentStatus)
	router.POST("/agent-fund", changeAgentFund)
	router.POST("/get-agent-simple-list", getAgentSimpleList)
	router.POST("/get-agent-fund-log", getAgentFundLog)
	router.POST("/get-agent-whitelist", getAgentWhitelist)
	router.POST("/get-agent-balance-snapshot", getAgentBalanceSnapshot)
	router.POST("/get-agent-ip-white-list", getAgentIpWhiteList)
	router.POST("/agent-ip-white-list", addAgentIpWhite)
	router.DELETE("/agent-ip-white-list", deleteAgentIpWhite)
	router.PUT("/agent-ip-white-list", updateAgentIpWhite)
	router.POST("/get-agent-product-list", getAgentProductList)
	router.POST("/agent-product", addAgentProduct)
	router.PUT("/agent-product", updateAgentProduct)
	router.DELETE("/agent-product", deleteAgentProduct)
	router.PUT("/batch-update-agent-product",batchUpdateAgentProduct)
	router.POST("/get-all-agent-product-list",getAllAgentProductList)
	router.PUT("/batch-update-agent-product-discount",batchUpdateAgentProductDiscount)
	router.POST("/get-agent-channel-list",getAgentChannelList)
	router.POST("/agent-channel",addAgentChannel)
	router.DELETE("/agent-channel",deleteAgentChannel)
	router.POST("/get-agent-product-channel-list",getAgentProductChannelList)
}

// 获取供货通道列表
func getAgentProductChannelList(c *gin.Context){
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"success":false,"message":"invalid input"})
		return
	}
	db := model.Db
	agent_id := requestData["agent_id"]
	switch agent_id.(type) {
	case float64:
		db = db.Where("agent_id = ?",agent_id)
	}
	agent_name := requestData["agent_name"]
	if agent_name != nil {
		agent_name2 := agent_name.(string)
		if agent_name2 != ""{
			db = db.Where("agent_name like ?","%"+agent_name2+"%")
		}
	}
	supplier_id := requestData["supplier_id"]
	switch supplier_id.(type) {
	case float64:
		db = db.Where("supplier_id = ?",supplier_id)
	}
	supplier_name := requestData["supplier_name"]
	if supplier_name != nil {
		supplier_name2 := supplier_name.(string)
		if supplier_name2 != ""{
			db = db.Where("supplier_name like ?","%"+supplier_name2+"%")
		}
	}
	business_type := requestData["business_type"]
	switch business_type.(type) {
	case float64:
		db = db.Where("business_type = ?",business_type)
	}
	product_id := requestData["product_id"]
	switch product_id.(type) {
	case float64:
		db = db.Where("product_id = ?",product_id)
	}
	product_category := requestData["product_category"]
	switch product_category.(type) {
	case float64:
		db = db.Where("product_category = ?",product_category)
	}
	operator := requestData["operator"]
	switch operator.(type) {
	case float64:
		db = db.Where("operator = ?",operator)
	}
	agent_product_id := requestData["agent_product_id"]
	switch agent_product_id.(type) {
	case float64:
		db = db.Where("agent_product_id = ?",agent_product_id)
	}
	supplier_product_id := requestData["supplier_product_id"]
	switch supplier_product_id.(type) {
	case float64:
		db = db.Where("supplier_product_id = ?",supplier_product_id)
	}
	up_product_id := requestData["up_product_id"].(string)
	if up_product_id != "" {
		db = db.Where("up_product_id = ?",up_product_id)
	}
	status := requestData["status"]
	switch status.(type) {
	case float64:
		db = db.Where("status = ?",status)
	}
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)

	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_product_channel
	db.Model(&model.Agent_product_channel{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK,gin.H{"success":true,"message":"","data":gin.H{
		"list":result,
		"total":total,
		"currentPage":page,
		"pageSize":pageSize,
	}})
}

// 添加代理商渠道
func addAgentChannel(c *gin.Context) {
	var agentChannel model.Agent_channel
	err := c.ShouldBindJSON(&agentChannel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Create(&agentChannel)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 删除代理商渠道
func deleteAgentChannel(c *gin.Context) {
	var ids []uint64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Agent_channel{})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 获取代理商渠道列表
func getAgentChannelList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	agent_id := requestData["agent_id"]
	switch agent_id.(type) {
	case float64:
		db = db.Where("agent_id = ?", agent_id)
	}
	supplier_id := requestData["supplier_id"]
	switch supplier_id.(type) {
	case float64:
		db = db.Where("supplier_id = ?", supplier_id)
	}
	supplier_name := requestData["supplier_name"]
	if supplier_name != nil {
		supplier_name2 := supplier_name.(string)
		if supplier_name2 != "" {
			db = db.Where("supplier_name like ?", "%"+supplier_name2+"%")
		}
	}
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_channel
	db.Model(&model.Agent_channel{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true, 
		"message": "", 
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		}})
}

// 批量修改代理商产品折扣
func batchUpdateAgentProductDiscount(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	product_list := requestData["product_list"].([]interface{})
	db := model.Db
	for _, v := range product_list {
		id := uint64(v.(map[string]interface{})["id"].(float64))
		discount := v.(map[string]interface{})["discount"].(float64)
		db.Model(&model.Agent_product{}).Where("id = ?", id).Update("discount", discount)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 获取所有代理商产品列表
func getAllAgentProductList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	db := model.Db
	var result []model.Agent_product
	db.Where("agent_id = ?",agent_id).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

// 批量修改代理商产品
func batchUpdateAgentProduct(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	//删除ids字段
	delete(requestData, "ids")
	//假如包括disabled_area，enabled_area，limit_operator字段，需要转换为字符串
	if _, ok := requestData["disabled_area"]; ok {
		requestData["disabled_area"] = convertArrayToString(requestData["disabled_area"].([]interface{}))
	}
	if _, ok := requestData["enabled_area"]; ok {
		requestData["enabled_area"] = convertArrayToString(requestData["enabled_area"].([]interface{}))
	}
	if _, ok := requestData["limit_operator"]; ok {
		requestData["limit_operator"] = convertArrayToString(requestData["limit_operator"].([]interface{}))
	}
	db := model.Db
	db.Model(&model.Agent_product{}).Where("id in ?", ids).Updates(requestData)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 删除代理商产品
func deleteAgentProduct(c *gin.Context) {
	var ids []uint64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Agent_product{})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 更新代理商产品
func updateAgentProduct(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	id := uint64(requestData["id"].(float64))
	db := model.Db
	var agentProduct model.Agent_product
	res := db.Where("id = ?", id).First(&agentProduct)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	agentProduct.Discount_type = uint8(requestData["discount_type"].(float64))
	agentProduct.Discount = requestData["discount"].(float64)
	agentProduct.Status = uint8(requestData["status"].(float64))
	agentProduct.Supply_strategy = uint8(requestData["supply_strategy"].(float64))
	agentProduct.Support_cache = uint8(requestData["support_cache"].(float64))
	agentProduct.Transfer_check = uint8(requestData["transfer_check"].(float64))
	agentProduct.Empty_check = uint8(requestData["empty_check"].(float64))
	agentProduct.Timeout_not_cache = uint8(requestData["timeout_not_cache"].(float64))
	agentProduct.Timeout = uint32(requestData["timeout"].(float64))
	agentProduct.Backup_channel_strategy = uint8(requestData["backup_channel_strategy"].(float64))
	agentProduct.Auto_submit_backup = uint8(requestData["auto_submit_backup"].(float64))
	agentProduct.Disabled_area = convertArrayToString(requestData["disabled_area"].([]interface{}))
	agentProduct.Enabled_area = convertArrayToString(requestData["enabled_area"].([]interface{}))
	agentProduct.Limit_operator = convertArrayToString(requestData["limit_operator"].([]interface{}))
	agentProduct.Remark = requestData["remark"].(string)
	db.Save(&agentProduct)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

func convertArrayToString(arr []interface{}) string {
	str := ""
	for _, v := range arr {
		str += fmt.Sprintf("%d", uint8(v.(float64)))
		str += ","
	}
	if len(str) > 0 {
		str = str[:len(str)-1]
	}
	return str
}

// 增加代理商产品
func addAgentProduct(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	agent_name := requestData["agent_name"].(string)
	product_list := requestData["product_list"].([]interface{})
	db := model.Db
	for _, v := range product_list {
		var agentProduct model.Agent_product
		agentProduct.Agent_id = uint64(agent_id)
		agentProduct.Agent_name = agent_name
		agentProduct.Business_type = uint8(v.(map[string]interface{})["type"].(float64))
		agentProduct.Product_id = uint64(v.(map[string]interface{})["id"].(float64))
		agentProduct.Product_name = v.(map[string]interface{})["name"].(string)
		agentProduct.Product_category = uint8(v.(map[string]interface{})["category"].(float64))
		agentProduct.Operator = uint8(v.(map[string]interface{})["operator"].(float64))
		agentProduct.Base_price = v.(map[string]interface{})["base_price"].(float64)
		agentProduct.Supply_strategy = 0
		agentProduct.Backup_channel_strategy = 0
		agentProduct.Discount_type = 0
		agentProduct.Discount = v.(map[string]interface{})["discount"].(float64)
		agentProduct.Timeout = 300
		agentProduct.Timeout_not_cache = 0
		agentProduct.Auto_submit_backup = 0
		agentProduct.Interal_time = 300
		agentProduct.Support_cache = 0
		agentProduct.Transfer_check = 0
		agentProduct.Empty_check = 0
		agentProduct.Disabled_area = v.(map[string]interface{})["disabled_area"].(string)
		agentProduct.Enabled_area = v.(map[string]interface{})["enabled_area"].(string)
		agentProduct.Limit_operator = v.(map[string]interface{})["limit_operator"].(string)
		agentProduct.Status = 0
		db.Create(&agentProduct)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 获取代理商产品列表
func getAgentProductList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_product
	db.Model(&model.Agent_product{}).Where("agent_id = ?", agent_id).Count(&total)
	db.Offset(int(offset)).Where("agent_id = ?", agent_id).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// 添加代理商白名单
func addAgentIpWhite(c *gin.Context) {
	var agentIpWhite model.Agent_whitelist
	err := c.ShouldBindJSON(&agentIpWhite)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agentIpWhite.Ip_location = middleware.IpLocation(agentIpWhite.Ip)
	db := model.Db
	db.Create(&agentIpWhite)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 删除代理商白名单
func deleteAgentIpWhite(c *gin.Context) {
	var ids []float64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Agent_whitelist{})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 更新代理商白名单
func updateAgentIpWhite(c *gin.Context) {
	var agentIpWhite model.Agent_whitelist
	err := c.ShouldBindJSON(&agentIpWhite)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agentIpWhite.Ip_location = middleware.IpLocation(agentIpWhite.Ip)
	db := model.Db
	db.Save(&agentIpWhite)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 获取代理商白名单列表
func getAgentIpWhiteList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"]
	ip := requestData["ip"].(string)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	switch agent_id.(type) {
	case float64:
		db = db.Where("agent_id = ?", agent_id)
	}
	if ip != "" {
		db = db.Where("ip like ?", "%"+ip+"%")
	}
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_whitelist
	db.Model(&model.Agent_whitelist{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// 获取代理商余额快照
func getAgentBalanceSnapshot(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_balance_snapshot
	db.Model(&model.Agent_balance_snapshot{}).Where("agent_id = ?", agent_id).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// 获取代理商白名单
func getAgentWhitelist(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_whitelist
	db.Model(&model.Agent_whitelist{}).Where("agent_id = ?", agent_id).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// 获取代理商资金操作日志
func getAgentFundLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_fund_log
	db.Model(&model.Agent_fund_log{}).Where("agent_id = ?", agent_id).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// 获取代理商简单列表
func getAgentSimpleList(c *gin.Context) {
	db := model.Db
	var result []map[string]interface{}
	db.Select("id", "name").Model(&model.Agent_account{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

// 操作代理商资金
func changeAgentFund(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent_id := requestData["agent_id"].(float64)
	agent_name := requestData["agent_name"].(string)
	fund_action := requestData["fund_action"].(string)
	amount := requestData["amount"].(float64)
	file := requestData["file"].(string)
	remark := requestData["remark"].(string)
	db := model.Db

	var agent model.Agent_account
	res := db.Where("id = ?", agent_id).First(&agent)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	beforeAmount := 0.0
	afterAmount := 0.0
	action := 0
	if fund_action == "subtract" || fund_action == "adjust" || fund_action == "add" {
		beforeAmount = agent.Fund_balance
		if fund_action == "add" {
			agent.Fund_balance += amount
			action = 0
		} else if fund_action == "subtract" {
			agent.Fund_balance -= amount
			action = 1
		} else if fund_action == "adjust" {
			agent.Fund_balance = amount
			action = 2
		}
		afterAmount = agent.Fund_balance
	} else {
		beforeAmount = agent.Credit_balance
		if fund_action == "credit_add" {
			agent.Credit_balance += amount
			action = 3
		} else if fund_action == "credit_subtract" {
			agent.Credit_balance -= amount
			action = 4
		}
		afterAmount = agent.Credit_balance
	}
	res = db.Save(&agent)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}

	// 添加资金变动记录
	var fundLog model.Agent_fund_log
	fundLog.Agent_id = uint64(agent_id)
	fundLog.Agent_name = agent_name
	fundLog.Action = uint8(action)
	fundLog.Amount = amount
	fundLog.Remark = remark
	fundLog.Cert_pic = file
	fundLog.Time = time.Now()
	fundLog.Before_amount = beforeAmount
	fundLog.After_amount = afterAmount
	res = db.Create(&fundLog)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "代理商资金变动成功"})
}

// 批量修改代理商状态
func batchUpdateAgentStatus(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"]
	status := requestData["status"]
	db := model.Db
	db.Model(&model.Agent_account{}).Where("id in ?", ids).Update("status", status)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 删除代理商
func deleteAgent(c *gin.Context) {
	var ids []float64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Agent_account{})
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 更新代理商
func updateAgent(c *gin.Context) {
	var agent model.Agent_account
	err := c.ShouldBindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Save(&agent)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 添加代理商
func addAgent(c *gin.Context) {
	var agent model.Agent_account
	err := c.ShouldBindJSON(&agent)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	agent.Id = 0
	agent.Secret_key = uuid.NewV4().String()
	agent.Secret_key = strings.Replace(agent.Secret_key, "-", "", -1)
	db := model.Db
	db.Create(&agent)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": ""})
}

// 获取代理商列表
func getAgentList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	id := requestData["id"]
	notification_method := requestData["notification_method"]
	status := requestData["status"]
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	switch id.(type) {
	case float64:
		db = db.Where("id = ?", id)
	}
	switch notification_method.(type) {
	case float64:
		db = db.Where("notification_method = ?", notification_method)
	}
	switch status.(type) {
	case float64:
		db = db.Where("status = ?", status)
	}
	offset := (page - 1) * pageSize
	var total int64
	var result []model.Agent_account
	db.Model(&model.Agent_account{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}
