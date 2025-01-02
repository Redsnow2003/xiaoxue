package module

import (
	"main/middleware"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//注册代理商接口路由
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
}

//添加代理商白名单
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

//删除代理商白名单
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

//更新代理商白名单
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

//获取代理商白名单列表
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

//获取代理商余额快照
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

//获取代理商白名单
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

//获取代理商资金操作日志
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

//获取代理商简单列表
func getAgentSimpleList(c *gin.Context) {
	db := model.Db
	var result []map[string]interface{}
	db.Select("id","name").Model(&model.Agent_account{}).Find(&result)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "", "data": result})
}

//操作代理商资金
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

//批量修改代理商状态
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

//删除代理商
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

//更新代理商
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

//添加代理商
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

//获取代理商列表
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