package model

// 产品分类表
type Product_category struct {
	Id            uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                // 分类id
	Category_name string `gorm:"column:category_name;type:varchar(255);not null" json:"category_name"` // 分类名称
}

// 表名
func (Product_category) TableName() string {
	return "product_category"
}

//产品信息表
type Product_information struct {
	Id                     uint64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`                        // 产品id
	Type                   uint8   `gorm:"column:type;type:tinyint" json:"type"`                                         // 业务类型
	Name                   string  `gorm:"column:name;type:varchar(255);not null" json:"name"`                           // 产品名称
	Category               uint64  `gorm:"column:category;type:int;not null" json:"category"`                            // 产品分类
	Operator               uint8   `gorm:"column:operator;type:tinyint" json:"operator"`                                 // 运营商
	Scope                  uint8   `gorm:"column:scope;type:tinyint" json:"scope"`                                       // 产品范围
	Price                  float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`                                 // 产品价格
	Unit                   uint8   `gorm:"column:unit;type:tinyint" json:"unit"`                                         // 产品单位
	Base_price             float64 `gorm:"column:base_price;type:decimal(10,2)" json:"base_price"`                       // 基础价格
	Disabled_area          string  `gorm:"column:disabled_area;type:varchar(255)" json:"disabled_area"`                  // 禁用地区
	Limit_operator         uint8   `gorm:"column:limit_operator;type:tinyint" json:"limit_operator"`                     // 限制运营商
	Sale_inventory         uint64  `gorm:"column:sale_inventory;type:bigint" json:"sale_inventory"`                      // 销售库存
	SSale_inventory_amount float64 `gorm:"column:sale_inventory_amount;type:decimal(10,2)" json:"sale_inventory_amount"` // 总库存
	Api_limit              uint32  `gorm:"column:api_limit;type:int" json:"api_limit"`                                   // API限制
	Remark                 string  `gorm:"column:remark;type:varchar(255)" json:"remark"`                                // 备注
}

// 表名
func (Product_information) TableName() string {
	return "product_information"
}

// product_update_log 产品更新日志表
type Product_update_log struct {
	ID                      uint64  `gorm:"column:id;primaryKey;autoIncrement;type:bigint unsigned" json:"id"`               // 自增主键
	Org_id                  uint64  `gorm:"column:org_id;type:bigint unsigned" json:"org_id"`                                // 组织ID
	User_id                 uint64  `gorm:"column:user_id;type:bigint unsigned" json:"user_id"`                              // 用户ID
	User_name               string  `gorm:"column:user_name;type:varchar(255)" json:"user_name"`                             // 用户名称
	User_type               uint8   `gorm:"column:user_type;type:tinyint" json:"user_type"`                                  // 用户类型
	Business_type           uint8   `gorm:"column:business_type;type:tinyint" json:"business_type"`                          // 业务类型
	Product_id              uint64  `gorm:"column:product_id;type:bigint unsigned" json:"product_id"`                        // 产品ID
	Product_name            string  `gorm:"column:product_name;type:varchar(255)" json:"product_name"`                       // 产品名称
	Product_category        uint64  `gorm:"column:product_category;type:bigint unsigned" json:"product_category"`            // 产品分类
	Operator                uint8   `gorm:"column:operator;type:tinyint" json:"operator"`                                    // 运营商
	Base_price              float64 `gorm:"column:base_price;type:decimal(10,2)" json:"base_price"`                          // 基础价格
	Supply_strategy         string  `gorm:"column:supply_strategy;type:varchar(255)" json:"supply_strategy"`                 // 供货策略
	Backup_channel_strategy string  `gorm:"column:backup_channel_strategy;type:varchar(255)" json:"backup_channel_strategy"` // 备用通道策略
	Discount_type           uint8   `gorm:"column:discount_type;type:tinyint" json:"discount_type"`                          // 折扣类型
	Discount                float64 `gorm:"column:discount;type:decimal(10,2)" json:"discount"`                              // 折扣
	Timeout                 uint32  `gorm:"column:timeout;type:int" json:"timeout"`                                          // 超时时长
	Timeout_not_cache       uint8   `gorm:"column:timeout_not_cache;type:tinyint" json:"timeout_not_cache"`                  // 超时不缓存
	Auto_submit_backup      uint8   `gorm:"column:auto_submit_backup;type:tinyint" json:"auto_submit_backup"`                // 自动提交备用
	Interal_time            uint32  `gorm:"column:interal_time;type:int" json:"interal_time"`                                // 间隔时间
	Support_cache           uint8   `gorm:"column:support_cache;type:tinyint" json:"support_cache"`                          // 支持缓存
	Transfer_check          uint8   `gorm:"column:transfer_check;type:tinyint" json:"transfer_check"`                        // 转移检查
	Empty_check             uint8   `gorm:"column:empty_check;type:tinyint" json:"empty_check"`                              // 空检查
	Disabled_area           string  `gorm:"column:disabled_area;type:varchar(255)" json:"disabled_area"`                     // 禁用地区
	Enabled_area            string  `gorm:"column:enabled_area;type:varchar(255)" json:"enabled_area"`                       // 启用地区
	Limit_operator          uint8   `gorm:"column:limit_operator;type:tinyint" json:"limit_operator"`                        // 限制运营商
	Status                  uint8   `gorm:"column:status;type:tinyint" json:"status"`                                        // 状态
	Executor_id             uint64  `gorm:"column:executor_id;type:bigint unsigned" json:"executor_id"`                      // 执行者ID
	Executor_name           string  `gorm:"column:executor_name;type:varchar(255)" json:"executor_name"`                     // 执行者名称
	Remark                  string  `gorm:"column:remark;type:varchar(255)" json:"remark"`                                   // 备注
}

// TableName 表名
func (Product_update_log) TableName() string {
	return "product_update_log"
}
