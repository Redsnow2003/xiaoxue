package model

// 用户表
type User struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 用户id
	UserName   			string `gorm:"column:username;type:varchar(255);not null" json:"username"` 	// 用户名
	Nickname   			string `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"` 	// 用户昵称
	Phone	  			string `gorm:"column:phone;type:varchar(255);not null" json:"phone"`       	// 用户电话
	Email	   			string `gorm:"column:email;type:varchar(255);not null" json:"email"`        // 用户邮箱
	Sex 	 			uint8  `gorm:"column:sex;type:tinyint;not null" json:"sex"` 				// 用户性别
	Status 				uint8  `gorm:"column:status;type:tinyint;not null" json:"status"` 			// 用户状态
	DeptId 				uint64 `gorm:"column:deptId;type:int;not null" json:"deptId"` 				// 部门id
	Password   			string `gorm:"column:password;type:varchar(255);not null" json:"password"` 	// 用户密码
	Roles      			string `gorm:"column:roles;type:int;not null" json:"roles"`               	// 用户角色
	Permissions        	string `gorm:"column:permissions;type:varchar(255)" json:"permissions"`     // 用户权限
	Avatar  			string `gorm:"column:avatar;type:datetime" json:"avatar"`          			// 头像
	CreateTime 			string `gorm:"column:createtime;type:datetime" json:"createtime"`      		// 创建时间
	Remark 				string `gorm:"column:remark;type:varchar(255)" json:"remark"`      			// 备注
}

// 表名
func (User) TableName() string {
	return "system_user"
}

// 菜单表
type Menu struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 菜单id
	ParentId   			uint64 `gorm:"column:parentId;type:int;not null" json:"parentId"`        	// 父级id
	MenuType   			uint8  `gorm:"column:menuType;type:tinyint;not null" json:"menuType"` // 菜单类型
	Title	  			string `gorm:"column:title;type:varchar(255);not null" json:"title"`       	// 菜单标题
	Name       			string `gorm:"column:name;type:varchar(255);not null" json:"name"`         	// 菜单名称
	Path       			string `gorm:"column:path;type:varchar(255);not null" json:"path"`         	// 菜单路径
	Component  			string `gorm:"column:component;type:varchar(255)" json:"component"`  // 菜单组件
	Rank 				int    `gorm:"column:rank;type:int" json:"rank"`                 	// 菜单排序
	Redirect  			string `gorm:"column:redirect;type:varchar(255)" json:"redirect"`          	// 菜单重定向
	Icon 	 			string `gorm:"column:icon;type:varchar(255)" json:"icon"`                 	// 菜单图标
	ExtraIcon 			string `gorm:"column:extraIcon;type:varchar(255)" json:"extraIcon"`     	// 菜单额外图标
	EnterTransition 	string `gorm:"column:enterTransition;type:varchar(255)" json:"enterTransition"` // 菜单进入动画
	LeaveTransition 	string `gorm:"column:leaveTransition;type:varchar(255)" json:"leaveTransition"` // 菜单离开动画
	ActivePath 			string `gorm:"column:leaveTransition;type:varchar(255)" json:"activePath"`      	// 菜单激活路径
	Roles 				string `gorm:"column:roles;type:varchar(255)" json:"roles"`       	// 菜单角色
	Auths 				string `gorm:"column:auths;type:varchar(255)" json:"auths"`       	// 菜单权限
	FrameSrc 			string `gorm:"column:frameSrc;type:varchar(255)" json:"frameSrc"`      	// 菜单内嵌地址
	FrameLoading 		bool   `gorm:"column:frameLoading;type:bool" json:"frameLoading"`      	// 菜单内嵌加载
	KeepAlive 			bool   `gorm:"column:keepAlive;type:bool" json:"keepAlive"`      	// 菜单缓存
	HiddenTag 			bool   `gorm:"column:hiddenTag;type:bool" json:"hiddenTag"`      	// 菜单隐藏标签
	FixedTag 			bool   `gorm:"column:fixedTag;type:bool" json:"fixedTag"`      	// 菜单固定标签
	ShowLink 			bool   `gorm:"column:showLink;type:bool" json:"showLink"`      	// 菜单显示链接
	ShowParent 			bool   `gorm:"column:showParent;type:bool" json:"showParent"`      	// 菜单显示父级
}

// 表名
func (Menu) TableName() string {
	return "system_menu"
}

// 部门表

type UpdateDeptData struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 部门id
	ParentId   			uint64 `gorm:"column:parentId;type:int;not null" json:"parentId"`        	// 父级id
	Name   				string `gorm:"column:name;type:varchar(255);not null" json:"name"` 			// 部门名称
	Sort 				uint8  `gorm:"column:sort;type:tinyint" json:"sort"`      					// 部门排序
	Phone  				string `gorm:"column:phone;type:varchar(255)" json:"phone"`      			// 部门电话
	Principal  			string `gorm:"column:principal;type:varchar(255)" json:"principal"`      	// 部门领导
	Email  				string `gorm:"column:email;type:varchar(255)" json:"email"`          		// 部门邮箱
	Status 				uint8  `gorm:"column:status;type:tinyint" json:"status"`      				// 部门状态
	Type 				uint8  `gorm:"column:type;type:tinyint" json:"type"`      					// 部门类型
	Remark 				string `gorm:"column:remark;type:varchar(255)" json:"remark"`      			// 备注
}

type Dept struct {
	UpdateDeptData
	CreateTime			string `gorm:"column:createTime;type:datetime" json:"createTime"`      		// 创建时间
}

// 表名

func (UpdateDeptData) TableName() string {
	return "system_dept"
}

func (Dept) TableName() string {
	return "system_dept"
}

// 角色表
type UpdateRoleData struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
	Name   				string `gorm:"column:name;type:varchar(255);not null" json:"name"` 			// 角色名称
	Code  				string `gorm:"column:code;type:varchar(255);not null" json:"code"`      	// 角色标识
	Status 				uint8  `gorm:"column:status;type:tinyint" json:"status"`      				// 角色状态
	Remark 				string `gorm:"column:remark;type:varchar(255)" json:"remark"`      			// 备注
	UpdateTime 			string `gorm:"column:updateTime;type:datetime" json:"updateTime"`      		// 更新时间
}

type Role struct {
	UpdateRoleData
	CreateTime 			string `gorm:"column:createTime;type:datetime" json:"createTime"`      		// 创建时间
}

// 表名
func (UpdateRoleData) TableName() string {
	return "system_role"
}

func (Role) TableName() string {
	return "system_role"
}

// 在线用户表
type OnlineUser struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 在线用户id
	Username   			string `gorm:"column:username;type:varchar(255);not null" json:"username"` 	// 用户名
	IP  				string `gorm:"column:ip;type:varchar(255);not null" json:"ip"` 				// 用户ip
	Address 			string `gorm:"column:address;type:varchar(255);not null" json:"address"` 	// 用户地址
	System 				string `gorm:"column:system;type:varchar(255);not null" json:"system"` 	// 用户系统
	Browser 			string `gorm:"column:browser;type:varchar(255);not null" json:"browser"` 	// 用户浏览器
	LoginTime 			string `gorm:"column:loginTime;type:datetime;not null" json:"loginTime"` 	// 登录时间
}

// 表名
func (OnlineUser) TableName() string {
	return "log_online"
}

// 登录日志表
type LoginLog struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 登录日志id
	Username   			string `gorm:"column:username;type:varchar(255);not null" json:"username"` 	// 用户名
	IP  				string `gorm:"column:ip;type:varchar(255);not null" json:"ip"` 				// 用户ip
	Address 			string `gorm:"column:address;type:varchar(255);not null" json:"address"` 	// 用户地址
	System 				string `gorm:"column:system;type:varchar(255);not null" json:"system"` 	// 用户系统
	Browser 			string `gorm:"column:browser;type:varchar(255);not null" json:"browser"` 	// 用户浏览器
	Status 				uint8  `gorm:"column:status;type:tinyint;not null" json:"status"` 			// 登录状态
	Behavior 			string `gorm:"column:behavior;type:varchar(255);not null" json:"behavior"` 	// 登录行为
	LoginTime 			string `gorm:"column:loginTime;type:datetime;not null" json:"loginTime"` 	// 登录时间
}

//表名
func (LoginLog) TableName() string {
	return "log_login"
}

// 操作日志表
type OperateLog struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 操作日志id
	Username   			string `gorm:"column:username;type:varchar(255);not null" json:"username"` 	// 用户名
	IP  				string `gorm:"column:ip;type:varchar(255);not null" json:"ip"` 				// 用户ip
	Address 			string `gorm:"column:address;type:varchar(255);not null" json:"address"` 	// 用户地址
	System 				string `gorm:"column:system;type:varchar(255);not null" json:"system"` 	// 用户系统
	Browser 			string `gorm:"column:browser;type:varchar(255);not null" json:"browser"` 	// 用户浏览器
	Status 				uint8  `gorm:"column:status;type:tinyint;not null" json:"status"` 			// 操作状态
	Summary 			string `gorm:"column:summary;type:varchar(255);not null" json:"summary"` 	// 操作摘要
	Module 				string `gorm:"column:module;type:varchar(255);not null" json:"module"` 	// 操作模块
	OperatingTime 		string `gorm:"column:operatingTime;type:datetime;not null" json:"operatingTime"` // 操作时间
}

// 表名
func (OperateLog) TableName() string {
	return "log_operation"
}

// 系统日志表
type SystemLog struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 系统日志id
	Level  				uint8  `gorm:"column:level;type:tinyint;not null" json:"level"` 			// 日志级别
	Module				string `gorm:"column:module;type:varchar(255);not null" json:"module"` 	// 日志模块
	Url  				string `gorm:"column:url;type:varchar(255);not null" json:"url"` 				// 请求路径
	Method 				string `gorm:"column:method;type:varchar(255);not null" json:"method"` 	// 请求方法
	Ip 					string `gorm:"column:ip;type:varchar(255);not null" json:"ip"` 	// 请求ip
	Address 			string `gorm:"column:address;type:varchar(255);not null" json:"address"` 	// 请求地址
	System 				string `gorm:"column:system;type:varchar(255);not null" json:"system"` 	// 请求系统
	Browser 			string `gorm:"column:browser;type:varchar(255);not null" json:"browser"` 	// 请求浏览器
	TakesTime 			int	   `gorm:"column:takesTime;type:int;not null" json:"takesTime"` 	// 请求耗时
	RequestTime 		string `gorm:"column:requestTime;type:datetime;not null" json:"requestTime"` // 请求时间
}

// 表名
func (SystemLog) TableName() string {
	return "log_system"
}
