package module

import (
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册供应商接口路由
func RegisterSupplierRoutes(router *gin.Engine) {
	router.POST("/get-supplier-template-list", getSupplierTemplateList)
	router.POST("/get-supplier-template-json", getSupplierTemplateJson)
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
	result := []model.Supply_template{}
	db.Model(&result).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message":"获取产品信息成功",
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
	var supplierTemplate model.Supply_template
	err := c.BindJSON(&supplierTemplate)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id = ?", supplierTemplate.Id).First(&supplierTemplate)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": supplierTemplate})
}