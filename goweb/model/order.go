package model

import "time"

// Order_list 订单表
type Order_list struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // 自增主键
	Business_type    uint8     `gorm:"column:business_type;type:tinyint" json:"business_type"`            // 业务类型
	Down_id          string    `gorm:"column:down_id;type:varchar(255)" json:"down_id"`                   // 第三方订单号
	Agent_id         uint64    `gorm:"column:agent_id;type:bigint unsigned" json:"agent_id"`              // 代理商ID
	Agent_name       string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`             // 代理商名称
	Product_category uint8     `gorm:"column:product_category;type:tinyint" json:"product_category"`      // 商品类别
	Product_id       uint64    `gorm:"column:product_id;type:bigint unsigned" json:"product_id"`          // 商品ID
	Product_name     string    `gorm:"column:product_name;type:varchar(255)" json:"product_name"`         // 商品名称
	Base_price       float64   `gorm:"column:base_price;type:double(10,2)" json:"base_price"`             // 基础价格
	Operator         uint8     `gorm:"column:operator;type:tinyint" json:"operator"`                      // 运营商类型
	Agent_discount   float64   `gorm:"column:agent_discount;type:double(10,2)" json:"agent_discount"`     // 代理商折扣
	Count            uint32    `gorm:"column:count;type:int" json:"count"`                                // 数量
	Recharge_number  string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`   // 充值号码
	Location         string    `gorm:"column:location;type:varchar(255)" json:"location"`                 // 所在地区
	Status           uint8     `gorm:"column:status;type:tinyint" json:"status"`                          // 状态
	Is_timeout       uint8     `gorm:"column:is_timeout;type:tinyint" json:"is_timeout"`                  // 是否超时
	Timeout          uint32    `gorm:"column:timeout;type:int" json:"timeout"`                            // 超时时长
	Is_cancel        uint8     `gorm:"column:is_cancel;type:tinyint" json:"is_cancel"`                    // 是否取消
	Create_time      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	Finish_time      time.Time `gorm:"column:finish_time;type:datetime" json:"finish_time"`               // 完成时间
	Notify_status    uint8     `gorm:"column:notify_status;type:tinyint" json:"notify_status"`            // 回调状态
	Special_params   string    `gorm:"column:special_params;type:varchar(255)" json:"special_params"`     // 特殊参数
	Remark           string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                     // 备注
}

// TableName 表名
func (Order_list) TableName() string {
	return "order_list"
}

// order_supplier 供应商订单表
type Order_supplier struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`   // 自增主键
	Order_id          uint64    `gorm:"column:order_id;type:bigint unsigned" json:"order_id"`                // 订单ID
	Business_type     uint8     `gorm:"column:business_type;type:tinyint" json:"business_type"`              // 业务类型
	Up_id             string    `gorm:"column:up_id;type:varchar(255)" json:"up_id"`                         // 上游订单号
	Agent_id          uint64    `gorm:"column:agent_id;type:bigint unsigned" json:"agent_id"`                // 代理商ID
	Agent_name        string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`               // 代理商名称
	Agent_discount    float64   `gorm:"column:agent_discount;type:double(10,2)" json:"agent_discount"`       // 代理商折扣
	Count             uint32    `gorm:"column:count;type:int" json:"count"`                                  // 数量
	Supplier_id       uint64    `gorm:"column:supplier_id;type:bigint unsigned" json:"supplier_id"`          // 供应商ID
	Supplier_name     string    `gorm:"column:supplier_name;type:varchar(255)" json:"supplier_name"`         // 供应商名称
	Product_id        uint64    `gorm:"column:product_id;type:bigint unsigned" json:"product_id"`            // 商品ID
	Product_name      string    `gorm:"column:product_name;type:varchar(255)" json:"product_name"`           // 商品名称
	Base_price        float64   `gorm:"column:base_price;type:double(10,2)" json:"base_price"`               // 基础价格
	Operator          uint8     `gorm:"column:operator;type:tinyint" json:"operator"`                        // 运营商
	Supplier_discount float64   `gorm:"column:supplier_discount;type:double(10,3)" json:"supplier_discount"` // 供应商折扣
	Recharge_number   string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`     // 充值号码
	Location          string    `gorm:"column:location;type:varchar(255)" json:"location"`                   // 所在地区
	Order_time        time.Time `gorm:"column:order_time;type:datetime" json:"order_time"`                   // 下单时间
	Status            uint8     `gorm:"column:status;type:tinyint" json:"status"`                            // 状态
	Create_time       time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                 // 创建时间
	Finish_time       time.Time `gorm:"column:finish_time;type:datetime" json:"finish_time"`                 // 完成时间
	Up_information    string    `gorm:"column:up_information;type:varchar(255)" json:"up_information"`       // 上游返回信息
	Update_time       time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                 // 更新时间
	Remark            string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                       // 备注
	Is_backup         uint8     `gorm:"column:is_backup;type:tinyint" json:"is_backup"`                      // 是否备份
}

// TableName 表名
func (Order_supplier) TableName() string {
	return "order_supplier"
}

// order_backup_submit_log 订单备用通道提交日志表
type Order_backup_submit_log struct {
	ID            uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`
	Order_id      uint64    `gorm:"column:order_id;type:bigint unsigned" json:"order_id"`
	Retry_count   uint32    `gorm:"column:retry_count;type:int" json:"retry_count"`
	Execute_count uint32    `gorm:"column:execute_count;type:int" json:"execute_count"`
	Interavl      uint32    `gorm:"column:interavl;type:int" json:"interavl"`
	Retry_time    time.Time `gorm:"column:retry_time;type:datetime" json:"retry_time"`
	Status        uint8     `gorm:"column:status;type:tinyint" json:"status"`
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
}

// TableName 表名
func (Order_backup_submit_log) TableName() string {
	return "order_backup_submit_log"
}

// order_submit_log 订单提交日志表
type Order_submit_log struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`  // 自增主键
	Order_id         uint64    `gorm:"column:order_id;type:bigint unsigned" json:"order_id"`               // 订单ID
	Request_header   string    `gorm:"column:request_header;type:varchar(1024)" json:"request_header"`     // 请求头信息
	Request_params   string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`     // 请求参数
	Request_ip       string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`              // 请求方IP
	Request_time     time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`              // 请求时间
	Response_context string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"` // 响应内容
	Response_time    time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`            // 响应时间
	Remark           string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                      // 备注
}

// TableName 表名
func (Order_submit_log) TableName() string {
	return "order_submit_log"
}

// order_query_log 订单查询日志表
type Order_query_log struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`  // 自增主键
	Order_id         uint64    `gorm:"column:order_id;type:bigint unsigned" json:"order_id"`               // 订单ID
	Request_header   string    `gorm:"column:request_header;type:varchar(1024)" json:"request_header"`     // 请求头信息
	Request_params   string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`     // 请求参数
	Request_ip       string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`              // 请求方IP
	Request_time     time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`              // 请求时间
	Response_context string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"` // 响应内容
	Response_time    time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`            // 响应时间
	Remark           string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                      // 备注
}

// TableName 表名
func (Order_query_log) TableName() string {
	return "order_query_log"
}

// order_notify_log 订单通知日志表
type Order_notify_log struct {
	ID            uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // 自增主键
	Order_id      uint64    `gorm:"column:order_id;type:bigint unsigned" json:"order_id"`              // 订单ID
	Send_log      string    `gorm:"column:send_log;type:varchar(1024)" json:"send_log"`                // 发送日志
	Send_address  string    `gorm:"column:send_address;type:varchar(255)" json:"send_address"`         // 发送地址
	Agent_return  string    `gorm:"column:agent_return;type:varchar(1024)" json:"agent_return"`        // 代理商返回信息
	Notify_status uint8     `gorm:"column:notify_status;type:tinyint" json:"notify_status"`            // 通知状态
	Notify_time   time.Time `gorm:"column:notify_time;type:datetime" json:"notify_time"`               // 通知时间
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                     // 备注
}

// TableName 表名
func (Order_notify_log) TableName() string {
	return "order_notify_log"
}

// order_supplier_submit_log 供货单提单日志表
type Order_supplier_submit_log struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`      // 自增主键
	Supplier_order_id uint64    `gorm:"column:supplier_order_id;type:bigint unsigned" json:"supplier_order_id"` // 供应商订单ID
	Request_header    string    `gorm:"column:request_header;type:varchar(1024)" json:"request_header"`         // 请求头信息
	Request_params    string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`         // 请求参数
	Request_ip        string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`                  // 请求方IP
	Request_time      time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`                  // 请求时间
	Response_context  string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"`     // 响应内容
	Response_time     time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`                // 响应时间
	Remark            string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 备注
}

// TableName 表名
func (Order_supplier_submit_log) TableName() string {
	return "order_supplier_submit_log"
}

// order_supplier_query_log 供货单查单日志表
type Order_supplier_query_log struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`      // 自增主键
	Supplier_order_id uint64    `gorm:"column:supplier_order_id;type:bigint unsigned" json:"supplier_order_id"` // 供应商订单ID
	Request_header    string    `gorm:"column:request_header;type:varchar(1024)" json:"request_header"`         // 请求头信息
	Request_params    string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`         // 请求参数
	Request_ip        string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`                  // 请求方IP
	Request_time      time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`                  // 请求时间
	Response_context  string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"`     // 响应内容
	Response_time     time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`                // 响应时间
	Remark            string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 备注
}

// TableName 表名
func (Order_supplier_query_log) TableName() string {
	return "order_supplier_query_log"
}

// order_supplier_cancel_log 供货单撤单日志表
type Order_supplier_cancel_log struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`      // 自增主键
	Supplier_order_id uint64    `gorm:"column:supplier_order_id;type:bigint unsigned" json:"supplier_order_id"` // 供应商订单ID
	Request_header    string    `gorm:"column:request_header;type:varchar(1024)" json:"request_header"`         // 请求头信息
	Request_params    string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`         // 请求参数
	Request_ip        string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`                  // 请求方IP
	Request_time      time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`                  // 请求时间
	Response_context  string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"`     // 响应内容
	Response_time     time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`                // 响应时间
	Remark            string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 备注
}

// TableName 表名
func (Order_supplier_cancel_log) TableName() string {
	return "order_supplier_cancel_log"
}

// order_supplier_callback_log 供货单回调日志表
type Order_supplier_callback_log struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`      // 自增主键
	Supplier_order_id uint64    `gorm:"column:supplier_order_id;type:bigint unsigned" json:"supplier_order_id"` // 供应商订单ID
	Callback_ip       string    `gorm:"column:callback_ip;type:varchar(255)" json:"callback_ip"`                // 回调IP
	Callback_context  string    `gorm:"column:callback_context;type:varchar(1024)" json:"callback_context"`     // 回调内容
	Callback_time     time.Time `gorm:"column:callback_time;type:datetime" json:"callback_time"`                // 回调时间
	Response_context  string    `gorm:"column:response_context;type:varchar(1024)" json:"response_context"`     // 响应内容
	Remark            string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 备注
}

// TableName 表名
func (Order_supplier_callback_log) TableName() string {
	return "order_supplier_callback_log"
}

// order_agent_intercept 代理商拦截订单表
type Order_agent_intercept struct {
	ID                   uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`          // 自增主键
	Down_id              string    `gorm:"column:down_id;type:varchar(255)" json:"down_id"`                            // 第三方订单号
	Agent_id             uint64    `gorm:"column:agent_id;type:bigint unsigned" json:"agent_id"`                       // 代理商ID
	Agent_name           string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`                      // 代理商名称
	Create_time          time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`                        // 创建时间
	Product_id           uint64    `gorm:"column:product_id;type:bigint unsigned" json:"product_id"`                   // 商品ID
	Product_name         string    `gorm:"column:product_name;type:varchar(255)" json:"product_name"`                  // 商品名称
	Recharge_number      string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`            // 充值号码
	Check_price          float64   `gorm:"column:check_price;type:double(10,2)" json:"check_price"`                    // 核对价格
	Request_ip           string    `gorm:"column:request_ip;type:varchar(255)" json:"request_ip"`                      // 请求方IP
	Request_time         time.Time `gorm:"column:request_time;type:datetime" json:"request_time"`                      // 请求时间
	Request_params       string    `gorm:"column:request_params;type:varchar(1024)" json:"request_params"`             // 请求参数
	Response_information string    `gorm:"column:response_information;type:varchar(1024)" json:"response_information"` // 响应信息
	Response_time        time.Time `gorm:"column:response_time;type:datetime" json:"response_time"`                    // 响应时间
	Remark               string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                              // 备注
}

// TableName 表名
func (Order_agent_intercept) TableName() string {
	return "order_agent_intercept"
}

// Order_Cache 订单表
type Order_cache struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // 自增主键
	Business_type    uint8     `gorm:"column:business_type;type:tinyint" json:"business_type"`            // 业务类型
	Down_id          string    `gorm:"column:down_id;type:varchar(255)" json:"down_id"`                   // 第三方订单号
	Agent_id         uint64    `gorm:"column:agent_id;type:bigint unsigned" json:"agent_id"`              // 代理商ID
	Agent_name       string    `gorm:"column:agent_name;type:varchar(255)" json:"agent_name"`             // 代理商名称
	Product_category uint8     `gorm:"column:product_category;type:tinyint" json:"product_category"`      // 商品类别
	Product_id       uint64    `gorm:"column:product_id;type:bigint unsigned" json:"product_id"`          // 商品ID
	Product_name     string    `gorm:"column:product_name;type:varchar(255)" json:"product_name"`         // 商品名称
	Base_price       float64   `gorm:"column:base_price;type:double(10,2)" json:"base_price"`             // 基础价格
	Operator         uint8     `gorm:"column:operator;type:tinyint" json:"operator"`                      // 运营商类型
	Agent_discount   float64   `gorm:"column:agent_discount;type:double(10,2)" json:"agent_discount"`     // 代理商折扣
	Count            uint32    `gorm:"column:count;type:int" json:"count"`                                // 数量
	Recharge_number  string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`   // 充值号码
	Location         string    `gorm:"column:location;type:varchar(255)" json:"location"`                 // 所在地区
	Status           uint8     `gorm:"column:status;type:tinyint" json:"status"`                          // 状态
	Is_timeout       uint8     `gorm:"column:is_timeout;type:tinyint" json:"is_timeout"`                  // 是否超时
	Timeout          uint32    `gorm:"column:timeout;type:int" json:"timeout"`                            // 超时时长
	Is_cancel        uint8     `gorm:"column:is_cancel;type:tinyint" json:"is_cancel"`                    // 是否取消
	Create_time      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	Finish_time      time.Time `gorm:"column:finish_time;type:datetime" json:"finish_time"`               // 完成时间
	Notify_status    uint8     `gorm:"column:notify_status;type:tinyint" json:"notify_status"`            // 回调状态
	Special_params   string    `gorm:"column:special_params;type:varchar(255)" json:"special_params"`     // 特殊参数
	Remark           string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                     // 备注
}

// TableName 表名
func (Order_cache) TableName() string {
	return "order_cache"
}

// 充值号码表
type Recharge_number struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // 自增主键
	Base_price       float64   `gorm:"column:base_price;type:double(10,2)" json:"base_price"`             // 基础价格
	Operator         uint8     `gorm:"column:operator;type:tinyint" json:"operator"`                      // 运营商类型
	Recharge_number  string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`   // 充值号码
	Location         string    `gorm:"column:location;type:varchar(255)" json:"location"`                 // 所在地区
	Status           uint8     `gorm:"column:status;type:tinyint" json:"status"`                          // 状态
	Create_time      time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	Finish_time      time.Time `gorm:"column:finish_time;type:datetime" json:"finish_time"`               // 完成时间
}

// TableName 表名
func (Recharge_number) TableName() string {
	return "order_list"
}

// order_number_blacklist 充值号码黑名单表
type Order_number_blacklist struct {
	ID               uint64    `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"` // 自增主键
	Recharge_number  string    `gorm:"column:recharge_number;type:varchar(255)" json:"recharge_number"`   // 充值号码
	Remark		   string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                     // 备注
}

// TableName 表名
func (Order_number_blacklist) TableName() string {
	return "order_number_blacklist"
}