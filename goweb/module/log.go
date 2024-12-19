package module

import (
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册日志路由
func RegisterLogRoutes(router *gin.Engine) {
	router.POST("/online-logs", getOnlineLogsList)
	router.POST("/login-logs", getLoginLogsList)
	router.DELETE("/login-logs", deleteLoginLogs)
	router.POST("/operation-logs", getOperationLogsList)
	router.DELETE("/operation-logs", deleteOperationLogs)
	router.POST("/system-logs", getSystemLogsList)
	router.DELETE("/system-logs", deleteSystemLogs)
}

// @Tags 日志
// @Summary 删除系统日志
func deleteSystemLogs(c *gin.Context) {
	deleteLogs(c, "log_system")
}

// @Tags 日志
// @Summary 获取系统日志
func getSystemLogsList(c *gin.Context) {
	getLogsList(c, &[]model.SystemLog{}, "requestTime")
}

// @Tags 日志
// @Summary 删除操作日志
func deleteOperationLogs(c *gin.Context) {
	deleteLogs(c, "log_operation")
}

// @Tags 日志
// @Summary 获取操作日志
func getOperationLogsList(c *gin.Context) {
	getLogsList(c, &[]model.OperateLog{}, "operatingTime")
}

// @Tags 日志
// @Summary 删除登录日志
func deleteLoginLogs(c *gin.Context) {
	deleteLogs(c, "log_login")
}

// @Tags 日志
// @Summary 获取登录日志
func getLoginLogsList(c *gin.Context) {
	getLogsList(c, &[]model.LoginLog{}, "loginTime")
}

// @Summary 获取在线用户日志
func getOnlineLogsList(c *gin.Context) {
	getLogsList(c, &[]model.OnlineUser{}, "")
}

// 删除日志的通用函数
func deleteLogs(c *gin.Context, tableName string) {
	model.Db.Exec("TRUNCATE TABLE " + tableName)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// 获取日志列表的通用函数
func getLogsList(c *gin.Context, logModel interface{}, timeField string) {
	var requestData map[string]interface{}
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	page := requestData["currentPage"].(float64)
	pageSize := requestData["pageSize"].(float64)
	db := model.Db
	offset := (page - 1) * pageSize

	// 根据请求数据构建查询条件
	for key, value := range requestData {
		switch key {
		case "module", "status", "username":
			if value != "" {
				db = db.Where(key+" like ?", "%"+value.(string)+"%")
			}
		case timeField:
			switch timeRange := value.(type) {
			case []interface{}:
				if len(timeRange) == 2 {
					db = db.Where(timeField+" between ? and ?", timeRange[0], timeRange[1])
				}
			}
		}
	}

	var total int64
	db.Model(logModel).Count(&total)
	db.Offset(int(offset)).Limit(int(pageSize)).Find(logModel)

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"list":        logModel,
			"total":       total,
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}
