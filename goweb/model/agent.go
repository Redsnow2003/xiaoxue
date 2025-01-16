package model

import "time"

// 代理商账户表
type Agent_account struct {
	Id                   uint64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                     // 代理商账户id
	Name                 string  `gorm:"column:name;type:varchar(255);not null" json:"name"`                        // 代理商名称
	Nickname             string  `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                         // 昵称
	Dept                 uint64  `gorm:"column:dept;type:int" json:"dept"`                                          // 部门
	Phone                string  `gorm:"column:phone;type:varchar(255)" json:"phone"`                               // 电话
	Email                string  `gorm:"column:email;type:varchar(255)" json:"email"`                               // 邮箱
	Secret_key           string  `gorm:"column:secret_key;type:varchar(255)" json:"secret_key"`                     // 密钥
	Notification_address string  `gorm:"column:notification_address;type:varchar(255)" json:"notification_address"` // 通知地址
	Notification_method  uint8   `gorm:"column:notification_method;type:tinyint" json:"notification_method"`        // 通知方式
	Customer             string  `gorm:"column:customer;type:varchar(255)" json:"customer"`                         // 客户
	Status               uint8   `gorm:"column:status;type:tinyint" json:"status"`                                  // 状态
	Fund_balance         float64 `gorm:"column:fund_balance;type:decimal(10,2)" json:"fund_balance"`                // 资金余额
	Credit_balance       float64 `gorm:"column:credit_balance;type:decimal(10,2)" json:"credit_balance"`            // 信用余额
	Frozen_amount        float64 `gorm:"column:frozen_amount;type:decimal(10,2)" json:"frozen_amount"`              // 冻结金额
	Cache_amount         float64 `gorm:"column:cache_amount;type:decimal(10,2)" json:"cache_amount"`                // 缓存金额
	Remark               string  `gorm:"column:remark;type:varchar(255)" json:"remark"`                             // 备注
}

// 表名
func (Agent_account) TableName() string {
	return "agent_account"
}

//代理商资金操作日志表
type Agent_fund_log struct {
	Id            uint64    `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`        // 代理商资金操作日志id
	Agent_id      uint64    `gorm:"column:agent_id;type:int" json:"agent_id"`                     // 代理商id
	Agent_name    string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`        // 代理商名称
	Action        uint8     `gorm:"column:action;type:tinyint" json:"action"`                     // 操作
	Time          time.Time `gorm:"column:time;type:datetime" json:"time"`                        // 时间
	Amount        float64   `gorm:"column:amount;type:decimal(10,2)" json:"amount"`               // 金额
	Before_amount float64   `gorm:"column:before_amount;type:decimal(10,2)" json:"before_amount"` // 操作前金额
	After_amount  float64   `gorm:"column:after_amount;type:decimal(10,2)" json:"after_amount"`   // 操作后金额
	Cert_pic      string    `gorm:"column:cert_pic;type:mediumtext" json:"cert_pic"`              // 凭证图片
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                // 备注
}

//表名
func (Agent_fund_log) TableName() string {
	return "agent_fund_log"
}

//代理商白名单表
type Agent_whitelist struct {
	Id          uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`   // 代理商白名单id
	Agent_id    uint64 `gorm:"column:agent_id;type:int" json:"agent_id"`                // 代理商id
	Agent_name  string `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`   // 代理商名称
	Ip          string `gorm:"column:ip;type:varchar(255)" json:"ip"`                   // 白名单ip
	Ip_location string `gorm:"column:ip_location;type:varchar(255)" json:"ip_location"` // ip归属地
	Remark      string `gorm:"column:remark;type:varchar(255)" json:"remark"`           // 备注
}

//表名
func (Agent_whitelist) TableName() string {
	return "agent_whitelist"
}

//代理商余额快照表
type Agent_balance_snapshot struct {
	Id             uint64    `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`          // 代理商余额快照id
	Agent_id       uint64    `gorm:"column:agent_id;type:int" json:"agent_id"`                       // 代理商id
	Agent_name     string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`          // 代理商名称
	Fund_balance   float64   `gorm:"column:fund_balance;type:decimal(10,2)" json:"fund_balance"`     // 资金余额
	Credit_balance float64   `gorm:"column:credit_balance;type:decimal(10,2)" json:"credit_balance"` // 信用余额
	Frozen_amount  float64   `gorm:"column:frozen_amount;type:decimal(10,2)" json:"frozen_amount"`   // 冻结金额
	Create_time    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`            // 创建时间
}

//表名
func (Agent_balance_snapshot) TableName() string {
	return "agent_balance_snapshot"
}

//代理商产品表
type Agent_product struct {
	Id                      uint64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                      // 代理商产品id
	Business_type           uint8   `gorm:"column:business_type;type:tinyint" json:"business_type"`                     // 业务类型
	Agent_id                uint64  `gorm:"column:agent_id;type:int" json:"agent_id"`                                   // 代理商id
	Agent_name              string  `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`                      // 代理商名称
	Product_id              uint64  `gorm:"column:product_id;type:int" json:"product_id"`                               // 产品id
	Product_name            string  `gorm:"column:product_name;type:varchar(255)" json:"product_name"`                  // 产品名称
	Product_category        uint8   `gorm:"column:product_category;type:tinyint" json:"product_category"`               // 产品类别
	Operator                uint8   `gorm:"column:operator;type:tinyint" json:"operator"`                               // 运营商
	Base_price              float64 `gorm:"column:base_price;type:decimal(10,2)" json:"base_price"`                     // 基础价格
	Supply_strategy         uint8   `gorm:"column:supply_strategy;type:tinyint" json:"supply_strategy"`                 // 供货策略
	Backup_channel_strategy uint8   `gorm:"column:backup_channel_strategy;type:tinyint" json:"backup_channel_strategy"` // 备用通道供货策略
	Discount_type           uint8   `gorm:"column:discount_type;type:tinyint" json:"discount_type"`                     // 折扣类型
	Discount                float64 `gorm:"column:discount;type:decimal(10,3)" json:"discount"`                         // 折扣
	Timeout                 uint32  `gorm:"column:timeout;type:int" json:"timeout"`                                     // 超时
	Timeout_not_cache       uint8   `gorm:"column:timeout_not_cache;type:tinyint" json:"timeout_not_cache"`             // 非缓存超时
	Auto_submit_backup      uint8   `gorm:"column:auto_submit_backup;type:tinyint" json:"auto_submit_backup"`           // 自动提交备份
	Interal_time            uint32  `gorm:"column:interal_time;type:int" json:"interal_time"`                           // 间隔时间
	Support_cache           uint8   `gorm:"column:support_cache;type:tinyint" json:"support_cache"`                     // 支持缓存
	Transfer_check          uint8   `gorm:"column:transfer_check;type:tinyint" json:"transfer_check"`                   // 转接检查
	Empty_check             uint8   `gorm:"column:empty_check;type:tinyint" json:"empty_check"`                         // 空号检查
	Disabled_area           string  `gorm:"column:disabled_area;type:varchar(255)" json:"disabled_area"`                // 禁用区域
	Enabled_area            string  `gorm:"column:enabled_area;type:varchar(255)" json:"enabled_area"`                  // 启用区域
	Limit_operator          string  `gorm:"column:limit_operator;type:varchar(255)" json:"limit_operator"`              // 限制运营商
	Status                  uint8   `gorm:"column:status;type:tinyint" json:"status"`                                   // 状态
	Remark                  string  `gorm:"column:remark;type:varchar(255)" json:"remark"`                              // 备注
}

//表名
func (Agent_product) TableName() string {
	return "agent_product"
}

//代理商通道表
type Agent_channel struct {
	Id            uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`       // 代理商通道id
	Agent_id      uint64 `gorm:"column:agent_id;type:int" json:"agent_id"`                    // 代理商id
	Agent_name    string `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`       // 代理商名称
	Supplier_id   uint64 `gorm:"column:supplier_id;type:int" json:"supplier_id"`              // 供应商id
	Supplier_name string `gorm:"column:supplier_name;type:varchar(255)" json:"supplier_name"` // 供应商名称4
}

//表名
func (Agent_channel) TableName() string {
	return "agent_channel"
}

//代理商产品通道表
type Agent_product_channel struct {
	Id                  uint64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                // 代理商产品通道id
	Business_type       uint8   `gorm:"column:business_type;type:tinyint" json:"business_type"`               // 业务类型
	Product_id          uint64  `gorm:"column:product_id;type:int" json:"product_id"`                         // 产品id
	Product_name        string  `gorm:"column:product_name;type:varchar(255)" json:"product_name"`            // 产品名称
	Product_category    uint8   `gorm:"column:product_category;type:tinyint" json:"product_category"`         // 产品类别
	Operator            uint8   `gorm:"column:operator;type:tinyint" json:"operator"`                         // 运营商
	Base_price          float64 `gorm:"column:base_price;type:decimal(10,2)" json:"base_price"`               // 基础价格
	Agent_id            uint64  `gorm:"column:agent_id;type:int" json:"agent_id"`                             // 代理商id
	Agent_name          string  `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`                // 代理商名称
	Agent_product_id    uint64  `gorm:"column:agent_product_id;type:int" json:"agent_product_id"`             // 代理商产品id
	Agent_discount      float64 `gorm:"column:agent_discount;type:decimal(10,3)" json:"agent_discount"`       // 代理商折扣
	Supply_strategy     uint8   `gorm:"column:supply_strategy;type:tinyint" json:"supply_strategy"`           // 供货策略
	Supplier_product_id uint64  `gorm:"column:supplier_product_id;type:int" json:"supplier_product_id"`       // 供货商产品id
	Supplier_id         uint64  `gorm:"column:supplier_id;type:int" json:"supplier_id"`                       // 供货商id
	Supplier_name       string  `gorm:"column:supplier_name;type:varchar(255)" json:"supplier_name"`          // 供货商名称
	Supplier_discount   float64 `gorm:"column:supplier_discount;type:decimal(10,3)" json:"supplier_discount"` // 供货商折扣
	Up_product_id       string  `gorm:"column:up_product_id;type:varchar(255)" json:"up_product_id"`          // 供货商名称
	Execute_count       uint8   `gorm:"column:execute_count;type:tinyint" json:"execute_count"`               // 执行次数
	Weight              uint8   `gorm:"column:weight;type:tinyint" json:"weight"`                             // 权重
	Priority            uint8   `gorm:"column:priority;type:tinyint" json:"priority"`                         // 优先级
	Province            uint8   `gorm:"column:province;type:tinyint" json:"province"`                         // 分省
	Status              uint8   `gorm:"column:status;type:tinyint" json:"status"`                             // 状态
	Is_connect          uint8   `gorm:"column:is_connect;type:tinyint" json:"is_connect"`                     // 是否可通
	Remark              string  `gorm:"column:remark;type:varchar(255)" json:"remark"`                        // 备注
}

//表名
//表名
func (Agent_product_channel) TableName() string {
	return "agent_product_channel"
}
