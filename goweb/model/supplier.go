package model

// 供应商模板表
type Supply_template struct {
	Id          					uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"` 						// 供应商模板id
	Name							string `gorm:"column:name;type:varchar(255);not null" json:"name"`   						// 模板名称
	Global_cb_address 				string `gorm:"column:global_cb_address;type:varchar(255)" json:"global_cb_address"` 		// 全局回调地址
	Is_need_product_id 				uint8  `gorm:"column:is_need_product_id;type:tinyint" json:"is_need_product_id"` 			// 是否需要产品id
	Is_bind_callback_address 		uint8  `gorm:"column:is_bind_callback_address;type:tinyint" json:"is_bind_callback_address"` // 是否绑定回调地址
	Is_support_inconsistent 		uint8  `gorm:"column:is_support_inconsistent;type:tinyint" json:"is_support_inconsistent"` 	// 是否支持不一致
	Is_support_cancel 				uint8  `gorm:"column:is_support_cancel;type:tinyint" json:"is_support_cancel"` 				// 是否支持取消
	Submit_address 					string `gorm:"column:submit_address;type:varchar(255)" json:"submit_address"` 				// 提交地址
	Query_address 					string `gorm:"column:query_address;type:varchar(255)" json:"query_address"` 				// 查询地址
	Balance_address 				string `gorm:"column:balance_address;type:varchar(255)" json:"balance_address"` 			// 结算地址
	Template_json 					string `gorm:"column:template_json;type:text" json:"template_json"` 						// 模板json
	Rremark 							string `gorm:"column:remark;type:varchar(255)" json:"remark"` 								// 备注
}

// 表名
func (Supply_template) TableName() string {
	return "supply_template"
}

//供应商账户表
type Supply_account struct {
	Id          					uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"` 					// 供应商账户id
	Name 							string `gorm:"column:name;type:varchar(255);not null" json:"name"` 						// 供应商名称
	Nickname 						string `gorm:"column:nickname;type:varchar(255)" json:"nickname"` 						// 昵称
	Dept 							uint64 `gorm:"column:dept;type:int" json:"dept"` 										// 部门
	Pohone 							string `gorm:"column:pohone;type:varchar(255)" json:"pohone"` 							// 电话
	Email 							string `gorm:"column:email;type:varchar(255)" json:"email"` 							// 邮箱
	Our_balance 					float64 `gorm:"column:our_balance;type:decimal(10,2)" json:"our_balance"` 				// 我方余额
	Up_balance 						float64 `gorm:"column:up_balance;type:decimal(10,2)" json:"up_balance"` 				// 上游余额
	Up_balance_update_time 			string `gorm:"column:up_balance_update_time;type:datetime" json:"up_balance_update_time"` 	// 上游余额更新时间
	Up_template 					string `gorm:"column:up_template;type:varchar(255)" json:"up_template"` 				// 上游模板
	Status 							uint8  `gorm:"column:status;type:tinyint" json:"status"` 								// 状态
	Status_info 					string `gorm:"column:status_info;type:varchar(255)" json:"status_info"` 				// 状态信息
	Remark							string `gorm:"column:remark;type:varchar(255)" json:"remark"` 							// 备注
}

// 表名
func (Supply_account) TableName() string {
	return "supply_account"
}


//供应商余额更新日志表
type Supply_balance_log struct {
	Id 							uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"` 					// 供应商余额更新日志id
	Supply_id 					uint64 `gorm:"column:supply_id;type:int" json:"supply_id"` 								// 供应商id
	Request_header 				string `gorm:"column:request_header;type:text" json:"request_header"` 					// 请求头
	Request_address 			string `gorm:"column:request_address;type:varchar(255)" json:"request_address"` 		// 请求地址
	Request_params 				string `gorm:"column:request_params;type:text" json:"request_params"` 					// 请求参数
	Request_time 				string `gorm:"column:request_time;type:datetime" json:"request_time"` 					// 请求时间
	Response_content 			string `gorm:"column:response_content;type:text" json:"response_content"` 				// 响应内容
	Response_time 				string `gorm:"column:response_time;type:datetime" json:"response_time"` 				// 响应时间
	Remark 						string `gorm:"column:remark;type:varchar(255)" json:"remark"` 							// 备注
}