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
	Disabled_area           string  `gorm:"column:disabled_area;type:varchar(255)" json:"disabled_area"`                    // 禁用地区
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
