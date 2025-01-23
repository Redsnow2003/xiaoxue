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
