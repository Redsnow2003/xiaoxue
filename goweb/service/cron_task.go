package service

import (
	"main/model"
	"time"

	"github.com/robfig/cron/v3"
)

//定时任务-执行代理商余额快照
func agentBalanceSnapshot() {
	db := model.Db
	var agents []model.Agent_account
	db.Find(&agents)
	for _, agent := range agents {
		var agentSnapshot model.Agent_balance_snapshot
		agentSnapshot.Agent_id = agent.Id
		agentSnapshot.Agent_name = agent.Name
		agentSnapshot.Fund_balance = agent.Fund_balance
		agentSnapshot.Credit_balance = agent.Credit_balance
		agentSnapshot.Frozen_amount = agent.Frozen_amount
		agentSnapshot.Create_time = time.Now()
		db.Create(&agentSnapshot)
	}
}

// StartCron 启动定时任务
func StartCron() {
	c := cron.New()
	c.AddFunc("00 00 * * *", agentBalanceSnapshot)
	c.Start()
}


