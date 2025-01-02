package model

import "time"

// 代理商账户表
type Agent_account struct {
	Id                     uint64    `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                     // 代理商账户id
	Name                   string    `gorm:"column:name;type:varchar(255);not null" json:"name"`                        // 代理商名称
	Nickname               string    `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                         // 昵称
	Dept				   uint64    `gorm:"column:dept;type:int" json:"dept"`                                          // 部门
	Phone                  string    `gorm:"column:phone;type:varchar(255)" json:"phone"`                               // 电话
	Email                  string    `gorm:"column:email;type:varchar(255)" json:"email"`                               // 邮箱
	Secret_key			   string    `gorm:"column:secret_key;type:varchar(255)" json:"secret_key"`                     // 密钥
	Notification_address   string    `gorm:"column:notification_address;type:varchar(255)" json:"notification_address"` // 通知地址
	Notification_method	   uint8     `gorm:"column:notification_method;type:tinyint" json:"notification_method"`        // 通知方式
	Customer 			   string    `gorm:"column:customer;type:varchar(255)" json:"customer"`                         // 客户
	Status				   uint8     `gorm:"column:status;type:tinyint" json:"status"`                                  // 状态
	Fund_balance		   float64   `gorm:"column:fund_balance;type:decimal(10,2)" json:"fund_balance"`                // 资金余额
	Credit_balance		   float64   `gorm:"column:credit_balance;type:decimal(10,2)" json:"credit_balance"`            // 信用余额
	Frozen_amount		   float64   `gorm:"column:frozen_amount;type:decimal(10,2)" json:"frozen_amount"`              // 冻结金额
	Cache_amount		   float64   `gorm:"column:cache_amount;type:decimal(10,2)" json:"cache_amount"`                // 缓存金额
	Remark				   string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                             // 备注
}

// 表名
func (Agent_account) TableName() string {
	return "agent_account"
}

//代理商资金操作日志表
type Agent_fund_log struct {
	Id               uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`           // 代理商资金操作日志id
	Agent_id         uint64 `gorm:"column:agent_id;type:int" json:"agent_id"`                        // 代理商id
	Agent_name	     string `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`           // 代理商名称
	Action		     uint8  `gorm:"column:action;type:tinyint" json:"action"`                        // 操作
	Time		     time.Time `gorm:"column:time;type:datetime" json:"time"`                            // 时间
	Amount		     float64 `gorm:"column:amount;type:decimal(10,2)" json:"amount"`                   // 金额
	Before_amount   float64 `gorm:"column:before_amount;type:decimal(10,2)" json:"before_amount"`     // 操作前金额
	After_amount	 float64 `gorm:"column:after_amount;type:decimal(10,2)" json:"after_amount"`       // 操作后金额
	Cert_pic		 string `gorm:"column:cert_pic;type:mediumtext" json:"cert_pic"`                    // 凭证图片
	Remark			 string `gorm:"column:remark;type:varchar(255)" json:"remark"`                     // 备注
}

//表名
func (Agent_fund_log) TableName() string {
	return "agent_fund_log"
}

//代理商白名单表
type Agent_whitelist struct {
	Id               uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`           // 代理商白名单id
	Agent_id         uint64 `gorm:"column:agent_id;type:int" json:"agent_id"`                        // 代理商id
	Agent_name	     string `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`           // 代理商名称
	Ip	             string `gorm:"column:ip;type:varchar(255)" json:"ip"`                			 // 白名单ip
	Ip_location		 string `gorm:"column:ip_location;type:varchar(255)" json:"ip_location"`         // ip归属地
	Remark			 string `gorm:"column:remark;type:varchar(255)" json:"remark"`                   // 备注
}

//表名
func (Agent_whitelist) TableName() string {
	return "agent_whitelist"
}

//代理商余额快照表
type Agent_balance_snapshot struct {
	Id               uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`           // 代理商余额快照id
	Agent_id         uint64 `gorm:"column:agent_id;type:int" json:"agent_id"`                        // 代理商id
	Agent_name	     string `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`           // 代理商名称
	Fund_balance	 float64 `gorm:"column:fund_balance;type:decimal(10,2)" json:"fund_balance"`     // 资金余额
	Credit_balance	 float64 `gorm:"column:credit_balance;type:decimal(10,2)" json:"credit_balance"` // 信用余额
	Frozen_amount	 float64 `gorm:"column:frozen_amount;type:decimal(10,2)" json:"frozen_amount"`   // 冻结金额
	Create_time		 time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`          // 创建时间
}

//表名
func (Agent_balance_snapshot) TableName() string {
	return "agent_balance_snapshot"
}