package model

// 用户表
type User struct {
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 用户id
	UserName   			string `gorm:"column:username;type:varchar(255);not null" json:"username"` 	// 用户名
	Password   			string `gorm:"column:password;type:varchar(255);not null" json:"password"` 	// 用户密码
	Nickname   			string `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"` 	// 用户昵称
	Roles      			string `gorm:"column:roles;type:int;not null" json:"roles"`               	// 用户角色
	Permissions        	string `gorm:"column:permissions;type:varchar(255)" json:"permissions"`     // 用户权限
	Avatar  			string `gorm:"column:avatar;type:datetime" json:"avatar"`          			// 头像
}

// 表名
func (User) TableName() string {
	return "users"
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
	ExtralIcon 			string `gorm:"column:extralIcon;type:varchar(255)" json:"extralIcon"`     	// 菜单额外图标
	EnterTransition 	string `gorm:"column:enterTransition;type:varchar(255)" json:"enterTransition"` // 菜单进入动画
	LeaveTransition 	string `gorm:"column:leave_transition;type:varchar(255)" json:"leaveTransition"` // 菜单离开动画
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
	return "menu"
}