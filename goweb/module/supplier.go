package module

import (
	"encoding/json"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 注册供应商接口路由
func RegisterSupplierRoutes(router *gin.Engine) {
	router.POST("/get-supplier-template-list", getSupplierTemplateList)
	router.POST("/get-supplier-template-json", getSupplierTemplateJson)
	router.POST("/get-supplier-simple-list", getSupplierSimpleList)
	router.POST("/get-supplier-template-name-list", getSupplierTemplateNameList)

	router.POST("/supplier-list", getSupplierList)
	router.POST("/supplier", addSupplier)
	router.PUT("/supplier", updateSupplier)
	router.DELETE("/supplier", deleteSupplier)
	router.PUT("/batch-update-supplier-status", batchUpdateSupplierStatus)
	router.POST("/supplier-change-fund", supplierChangeFund)
	router.POST("/get-supplier-fund-log", getSupplierFundLog)
	router.POST("/get-supplier-product", getSupplierProductList)
	router.POST("/get-supplier-up-balance-log", getSupplierUpBalanceLog)
}

// @Tags 供应商
// @Summary 获取供应商上游余额更新日志
func getSupplierUpBalanceLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	supplierId := requestData["supplier_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	result := []model.Supplier_balance_log{}
	db.Model(&result).Where("supplier_id = ?", supplierId).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Where("supplier_id = ?", supplierId).Find(&result)
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

// @Tags 供应商
// @Summary 获取供应商产品列表
func getSupplierProductList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	business_type := requestData["business_type"]
	supplier_id := requestData["supplier_id"]
	product_category := requestData["product_category"]
	product_id := requestData["product_id"]
	product_name := requestData["product_name"].(string)
	operator := requestData["operator"]
	up_product_id := requestData["up_product_id"].(string)
	status := requestData["status"]
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize

	switch business_type.(type) {
	case float64:
		db = db.Where("business_type = ?", business_type)
	}
	switch supplier_id.(type) {
	case float64:
		db = db.Where("supplier_id = ?", supplier_id)
	}
	switch product_category.(type) {
	case float64:
		db = db.Where("product_category = ?", product_category)
	}
	switch product_id.(type) {
	case float64:
		db = db.Where("product_id = ?", product_id)
	}
	if product_name != "" {
		db = db.Where("up_template liek ?", "%"+product_name+"%")
	}
	switch operator.(type) {
	case float64:
		db = db.Where("operator = ?", operator)
	}
	if up_product_id != "" {
		db = db.Where("up_template liek ?", "%"+up_product_id+"%")
	}
	switch status.(type) {
	case float64:
		db = db.Where("status = ?", status)
	}

	var total int64
	result := []model.Supplier_product{}
	db.Model(&result).Count(&total)
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

// @Tags 供应商
// @Summary 获取供应商资金操作日志
func getSupplierFundLog(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	supplierId := requestData["supplier_id"].(float64)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	var total int64
	result := []model.Supplier_fund_log{}
	db.Model(&result).Where("supplier_id = ?", supplierId).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Where("supplier_id = ?", supplierId).Find(&result)
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

// @Tags 供应商
// @Summary 供应商资金变动
func supplierChangeFund(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	supplierId := requestData["supplier_id"].(float64)
	supplierName := requestData["supplier_name"].(string)
	fundAction := requestData["fund_action"].(string)
	amount := requestData["amount"].(float64)
	file := requestData["file"].(string)
	remark := requestData["remark"].(string)
	db := model.Db
	// 对标supplier_account表的balance字段进行加减
	var supplier model.Supplier_account
	res := db.Where("id = ?", supplierId).First(&supplier)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}

	beforeAmount := supplier.Our_balance
	afterAmount := 0.0
	action := 0
	if fundAction == "add" {
		supplier.Our_balance += amount
		action = 0
	} else if fundAction == "subtract" {
		supplier.Our_balance -= amount
		action = 1
	} else if fundAction == "adjust" {
		supplier.Our_balance = amount
		action = 2
	}
	afterAmount = supplier.Our_balance
	res = db.Save(&supplier)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	// 添加资金变动记录
	var fundLog model.Supplier_fund_log
	fundLog.Supplier_id = uint64(supplierId)
	fundLog.Supplier_name = supplierName
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
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "供应商资金变动成功"})
}

// @Tags 供应商
// @Summary 批量更新供应商状态
func batchUpdateSupplierStatus(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	ids := requestData["ids"].([]interface{})
	status := requestData["status"].(float64)
	db := model.Db
	res := db.Model(&model.Supplier_account{}).Where("id in ?", ids).Update("status", status)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "批量更新供应商状态成功"})
}

// @Tags 供应商
// @Summary 更新供应商
func updateSupplier(c *gin.Context) {
	var supplier model.Supplier_account
	err := c.BindJSON(&supplier)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	if supplier.Nickname == "" {
		supplier.Nickname = supplier.Name
	}
	db := model.Db
	// 排除up_balance_update_time字段
	res := db.Omit("up_balance_update_time").Save(&supplier)
	// 返回添加结果
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "更新供应商成功"})
}

// @Tags 供应商
// @Summary 删除供应商
func deleteSupplier(c *gin.Context) {
	var ids []float64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	res := db.Delete(&model.Supplier_account{}, ids)
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除供应商成功"})
}

// @Tags 供应商
// @Summary 添加供应商
func addSupplier(c *gin.Context) {
	var supplier model.Supplier_account
	err := c.BindJSON(&supplier)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	if supplier.Nickname == "" {
		supplier.Nickname = supplier.Name
	}
	db := model.Db
	// 排除up_balance_update_time字段
	res := db.Omit("up_balance_update_time").Create(&supplier)
	// 返回添加结果
	if res.Error != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": res.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "添加供应商成功"})
}

// @Tags 供应商
// @Summary 获取供应商列表
func getSupplierList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	name := requestData["name"]
	up_template := requestData["up_template"].(string)
	status := requestData["status"]
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	switch name.(type) {
	case float64:
		db = db.Where("id = ?", name)
	}
	if up_template != "" {
		db = db.Where("up_template = ?", up_template)
	}
	switch status.(type) {
	case float64:
		db = db.Where("status = ?", status)
	}
	var total int64
	result := []model.Supplier_account{}
	db.Model(&result).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取产品信息成功",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// @Tags 供应商
// @Summary 获取供应商模板名称列表
func getSupplierTemplateNameList(c *gin.Context) {
	db := model.Db
	var result []string
	db.Model(&model.Supplier_template{}).Pluck("name", &result)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

// @Tags 供应商
// @Summary 获取供应商简单列表
func getSupplierSimpleList(c *gin.Context) {
	db := model.Db
	var result []model.Supplier_account_simple
	db.Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

// @Tags 供应商
// @Summary 获取供应商模板列表
func getSupplierTemplateList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	name := requestData["name"].(string)
	submit_address := requestData["submit_address"].(string)
	query_address := requestData["query_address"].(string)
	balance_address := requestData["balance_address"].(string)
	remark := requestData["remark"]
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	if submit_address != "" {
		db = db.Where("submit_address like ?", "%"+submit_address+"%")
	}
	if query_address != "" {
		db = db.Where("query_address like ?", "%"+query_address+"%")
	}
	if balance_address != "" {
		db = db.Where("balance_address like ?", "%"+balance_address+"%")
	}
	if remark != "" {
		db = db.Where("remark like ?", "%"+remark.(string)+"%")
	}
	var total int64
	result := []model.Supplier_template{}
	db.Model(&result).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取产品信息成功",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// @Tags 供应商
// @Summary 获取供应商模板json
func getSupplierTemplateJson(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	template_name := requestData["template_name"].(string)
	db := model.Db
	var supplierTemplate model.Supplier_template
	db.Select("template_json").Where("name = ?", template_name).First(&supplierTemplate)
	var jsonObj map[string]interface{}
	err = json.Unmarshal([]byte(supplierTemplate.Template_json), &jsonObj)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": jsonObj})
}
