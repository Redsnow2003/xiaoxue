package module

import (
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册产品接口路由
func RegisterProductRoutes(router *gin.Engine) {
	router.POST("/get-product-category", getProductCategoryList)
	router.POST("/product-category", addProductCategory)
	router.PUT("/product-category", updateProductCategory)
	router.DELETE("/product-category", deleteProductCategory)

	router.POST("/get-product-list", getProductInformationList)
	router.POST("/product-info", addProductInformation)
	router.PUT("/product-info", updateProductInformation)
	router.DELETE("/product-info", deleteProductInformation)
	router.POST("/product-info/import", importProductInformation)
	router.POST("/product-info/export", exportProductInformation)
	router.POST("/get-product-id-name",getProductInformationIdAndName)
}

// @Tags 产品
// @Summary 获取产品id和name
func getProductInformationIdAndName(c *gin.Context) {
	db := model.Db
	var result []model.Product_information
	db.Select("id, name").Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

// @Tags 产品
// @Summary 更新产品信息
func updateProductInformation(c *gin.Context) {
	var productInformation model.Product_information
	err := c.BindJSON(&productInformation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Save(&productInformation)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": productInformation})
}

// @Tags 产品
// @Summary 删除产品信息
func deleteProductInformation(c *gin.Context) {
	var ids []uint64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Product_information{})
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Tags 产品
// @Summary 增加产品信息
func addProductInformation(c *gin.Context) {
	var productInformation model.Product_information
	err := c.BindJSON(&productInformation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Create(&productInformation)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": productInformation})
}

// @Tags 产品
// @Summary 导入产品信息
func importProductInformation(c *gin.Context) {
	var requestData []model.Product_information
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Create(&requestData)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Tags 产品
// @Summary 获取产品信息
func getProductInformationList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	id := requestData["id"].(string)
	type1 := requestData["type"]
	name := requestData["name"].(string)
	category := requestData["category"]
	operator := requestData["operator"]
	price := requestData["price"].(string)
	scope := requestData["scope"]
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	if id != "" {
		db = db.Where("id = ?", id)
	}
	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}
	switch type1.(type) {
	case float64:
		db = db.Where("type = ?", type1)
	}
	switch category.(type) {
	case float64:
		db = db.Where("category = ?", category)
	}
	switch operator.(type) {
	case float64:
		db = db.Where("operator = ?", operator)
	}
	if price != "" {
		db = db.Where("price = ?", price)
	}
	switch scope.(type) {
	case float64:
		db = db.Where("scope = ?", scope)
	}
	var total int64
	result := []model.Product_information{}
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

// @Tags 产品
// @Summary 导出产品信息
func exportProductInformation(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var result []model.Product_information
	db.Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

// @Tags 产品
// @Summary 增加产品分类
func addProductCategory(c *gin.Context) {
	var productCategory model.Product_category
	err := c.BindJSON(&productCategory)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Create(&productCategory)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": productCategory})
}

// @Tags 产品
// @Summary 更新产品分类
func updateProductCategory(c *gin.Context) {
	var productCategory model.Product_category
	err := c.BindJSON(&productCategory)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Save(&productCategory)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": productCategory})
}

// @Tags 产品
// @Summary 删除产品分类
func deleteProductCategory(c *gin.Context) {
	var ids []uint64
	err := c.BindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	db.Where("id in ?", ids).Delete(&model.Product_category{})
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// @Tags 产品
// @Summary 获取产品分类
func getProductCategoryList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	category := requestData["category_name"].(string)
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize
	if category != "" {
		db = db.Where("category_name like ?", "%"+category+"%")
	}
	var total int64
	result := []model.Product_category{}
	db.Model(&result).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)

		// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message":"获取产品分类成功",
		"data": gin.H{
			"list":        result,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}