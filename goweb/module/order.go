package module

import (
	"main/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 注册订单路由
func RegisterOrderRoutes(router *gin.Engine) {
	router.POST("/order-list", getOrderList)
}

// @Tags 订单
// @Summary 获取订单列表
func getOrderList(c *gin.Context) {
	var requestData map[string]interface{}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"success":false,"message":"invalid input"})
		return
	}
	db := model.Db
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)

	offset := (page - 1) * pageSize
	var total int64
	var result []model.Order_list
	db.Model(&model.Order_list{}).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(&result)
	c.JSON(http.StatusOK,gin.H{"success":true,"message":"","data":gin.H{
		"list":result,
		"total":total,
		"currentPage":page,
		"pageSize":pageSize,
	}})
}