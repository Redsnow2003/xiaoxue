package module

import (
	"main/common"
	"main/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterDepartmentRoutes(router *gin.Engine) {
	router.GET("/dept", getDept)
	router.POST("/dept", addDept)
	router.PUT("/dept", updateDept)
	router.DELETE("/dept", deleteDept)
}

// @Tags 部门
// @Summary 获取部门列表
// @Description 获取所有部门的列表
// @Produce json
// @Success 200 {object} gin.H{"success": bool, "data": []model.Dept}
// @Failure 200 {object} gin.H{"success": bool, "message": string}
// @Router /dept [get]
func getDept(c *gin.Context) {
	db := model.Db
	depts := []model.Dept{}
	if err := db.Find(&depts).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "failed to get departments"})
		return
	}
	// Implement logic to get departments
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    depts,
		})
}

func addDept(c *gin.Context) {
	var deptVar model.Dept
	err := c.ShouldBindJSON(&deptVar)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	parentId := deptVar.ParentId
	db := model.Db
	var maxId int
	if parentId != 0 {
		// check parent department
		var dept model.Dept
		res := db.Where("id = ?", parentId).First(&dept)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "parent department not found"})
			return
		}
		//
		if err := db.Table("system_dept").Where("parentId = ?", parentId).Select("MAX(id)").Scan(&maxId).Error; err != nil {
			maxId = int(parentId + 1)
		} else {
			maxId = maxId + 1
		}
	} else {
		if err := db.Table("system_dept").Where("parentId = 0").Select("MAX(id)").Scan(&maxId).Error; err != nil {
			maxId = 100
		} else {
			maxId = maxId + 100
		}
	}

	deptVar.Id = uint64(maxId)
	deptVar.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	db.Create(&deptVar)
	// Implement logic to create a department
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "create department"})
}

func updateDept(c *gin.Context) {
	var deptVar model.UpdateDeptData
	err := c.ShouldBindJSON(&deptVar)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	db.Save(&deptVar)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update department"})
}

func deleteDept(c *gin.Context) {
	var deptVar model.UpdateDeptData
	err := c.ShouldBindJSON(&deptVar)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	//先删除子部门
	db.Where("parentId = ?", deptVar.Id).Delete(&model.Dept{})
	// Implement logic to delete a department
	db.Delete(&deptVar)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete department"})
}

// 注册菜单路由
func RegisterMenuRoutes(router *gin.Engine) {
	routerName := "/menu"
	router.GET(routerName, getMenu)
	router.POST(routerName, addMenu)
	router.PUT(routerName, updateMenu)
	router.DELETE(routerName, deleteMenu)
}

//获取菜单信息
func getMenu(c *gin.Context) {
	menus := []model.Menu{}
	db := model.Db
	db.Find(&menus)
	// Implement logic to get departments
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    menus,
		})
}

//添加菜单信息
func addMenu(c *gin.Context) {
	var menu model.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	parentId := menu.ParentId
	db := model.Db
	var maxId int
	if parentId != 0 {
		// check parent department
		var menu2 model.Menu
		res := db.Where("id = ?", parentId).First(&menu2)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "parent menu not found"})
			return
		}
		//
		if err := db.Table("system_menu").Where("parentId = ?", parentId).Select("MAX(id)").Scan(&maxId).Error; err != nil {
			maxId = int(parentId + 1)
		} else {
			maxId = maxId + 1
		}
	} else {
		if err := db.Table("system_menu").Where("parentId = 0").Select("MAX(id)").Scan(&maxId).Error; err != nil {
			maxId = 100
		} else {
			maxId = maxId + 100
		}
	}
	menu.Id = uint64(maxId)
	if db.Create(&menu) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "create menu failed"})
		return
	}
	// Implement logic to create a department
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "create menu"})
}

//更新菜单信息
func updateMenu(c *gin.Context) {
	var menu model.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	if db.Save(&menu) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update menu failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update menu"})
}

//删除菜单信息
func deleteMenu(c *gin.Context) {
	var menu model.Menu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to delete a department
	db.Where("parentId = ?", menu.Id).Delete(&model.Menu{})
	if db.Delete(&menu) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "delete menu failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete menu"})
}

// 注册角色路由
func RegisterRoleRoutes(router *gin.Engine) {
	routerName := "/role"
	router.GET(routerName, getRole)
	router.POST(routerName, addRole)
	router.PUT(routerName, updateRole)
	router.DELETE(routerName, deleteRole)
	router.GET("/list-all-role",getAllRoleList)
	router.POST("/list-role-ids",getRoleIds)
	router.POST("/role-menu",getRoleMenu)
	router.POST("/role-menu-ids",getRoleMenuIds)
	router.POST("/update-role-menu",updateRoleMenuIds)
	router.PUT("/role-status",updateRoleStatus)
}
type statusData struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
	Status 				uint8  `gorm:"column:status;type:tinyint" json:"status"`      				// 角色状态
}
func (statusData) TableName() string {
	return "system_role"
}
//更新角色状态
func updateRoleStatus(c *gin.Context) {
	var role statusData
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	if db.Save(&role) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update role failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update role"})
}

type roleMenu struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 菜单id
	ParentId   			uint64 `gorm:"column:parentId;type:int;not null" json:"parentId"`        	// 父级id
	MenuType   			uint8  `gorm:"column:menuType;type:tinyint;not null" json:"menuType"` // 菜单类型
	Title	  			string `gorm:"column:title;type:varchar(255);not null" json:"title"`       	// 菜单标题
}

// 表名
func (roleMenu) TableName() string {
	return "system_menu"
}

//获取角色菜单
func getRoleMenu(c *gin.Context) {
	db := model.Db
	roleMenus := []roleMenu{}
	db.Find(&roleMenus)
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    roleMenus,
		})
}

type roleMenuData struct{
	Id 		uint64 	`json:"id"`  // 角色id
	Menus []int 	`json:"menus"`  // 菜单id
}

//更新角色菜单ids
func updateRoleMenuIds(c *gin.Context) {
	var roleMenuData roleMenuData
	err := c.ShouldBindJSON(&roleMenuData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var role model.Role
	res := db.Where("id = ?", roleMenuData.Id).First(&role)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "role not found"})
		return
	}

	menus := []model.Menu{}
	db.Find(&menus)
	for _, menu := range menus {
		menuRoles := strings.Split(menu.Roles, ",")
		if common.IsContain(roleMenuData.Menus, int(menu.Id)) {
			if menu.Roles == "" {
				menu.Roles = role.Code
				db.Save(&menu)
			} else {
				if !common.IsContain(menuRoles, role.Code) {
					menu.Roles = menu.Roles + "," + role.Code
					db.Save(&menu)
				}
			}	
		} else {
			if common.IsContain(menuRoles, role.Code) {
				menuRoles = common.RemoveStr(menuRoles, role.Code)
				menu.Roles = strings.Join(menuRoles, ",")
				db.Save(&menu)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update role menu ids"})
}

type idData struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
}
//获取角色菜单ids
func getRoleMenuIds(c *gin.Context) {
	var idData idData
	err := c.ShouldBindJSON(&idData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var role model.Role
	res := db.Where("id = ?", idData.Id).First(&role)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "role not found"})
		return
	}
	roleMenuIds := make([]uint64, 0)
	menus := []model.Menu{}
	db.Find(&menus)
	for _, menu := range menus {
		menuRoles := strings.Split(menu.Roles, ",")
		if common.IsContain(menuRoles, role.Code) {
			roleMenuIds = append(roleMenuIds, menu.Id)
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": roleMenuIds})
}

type roleData struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
	Name   				string `gorm:"column:name;type:varchar(255);not null" json:"name"` 			// 角色名称
}

// 表名
func (roleData) TableName() string {
	return "system_role"
}

//获取所有角色信息
func getAllRoleList(c *gin.Context) {
	roles := []roleData{}
	db := model.Db
	db.Find(&roles)
	// Implement logic to get departments
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data":    roles,
		})
}

type roleIds struct{
	UserId 	int	`json:"userId"`  // 用户id
}
//获取角色ids
func getRoleIds(c *gin.Context) {
	var roleIds roleIds
	err := c.ShouldBindJSON(&roleIds)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var user model.User
	res := db.Where("id = ?", roleIds.UserId).First(&user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "user not found"})
		return
	}
	roles := strings.Split(user.Roles,",")
	roleids := make([]uint64, 0)
	for _, role := range roles {
		var _role model.Role
		res = db.Where("code = ?", role).First(&_role)
		if res.RowsAffected == 0 {
			continue
		}
		roleids = append(roleids, _role.Id)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": roleids})
}

//获取角色信息
func getRole(c *gin.Context) {
	roles := []model.Role{}
	db := model.Db
	db.Find(&roles)
	// Implement logic to get departments
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data": gin.H{
				"list": roles,
				"total": len(roles), // 总条目数
          		"pageSize": 10, // 每页显示条目个数
          		"currentPage": 1, // 当前页数
			},
		})
}

//添加角色信息
func addRole(c *gin.Context) {
	var role model.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	role.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	role.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	db := model.Db
	if db.Create(&role) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "create role failed"})
		return
	}
	// Implement logic to create a department
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "create role"})
}

//更新角色信息
func updateRole(c *gin.Context) {
	var role model.UpdateRoleData
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var _role model.Role
	db.Where("id = ?", role.Id).First(&_role)
	//假如角色code发生变化，需要更新菜单表和用户表中的角色字段
	if _role.Code != role.Code {
		updateRoleMenuAndUser(_role.Code, role.Code)
	}
	role.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	// Implement logic to update a department
	if db.Save(&role) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update role failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update role"})
}

//删除角色信息
func deleteRole(c *gin.Context) {
	var role model.Role
	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	var _role model.Role
	db.Where("id = ?", role.Id).First(&_role)
	deleteRoleMenuAndUser(_role.Code)
	// Implement logic to delete a department
	db.Delete(&role)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete role"})
}

// 更新菜单表和用户表中角色字段
func updateRoleMenuAndUser(oldCode string, newCode string) {
	db := model.Db
	menus := []model.Menu{}
	db.Find(&menus)
	for _, menu := range menus {
		menuRoles := strings.Split(menu.Roles, ",")
		if common.IsContain(menuRoles, oldCode) {
			menuRoles = common.RemoveStr(menuRoles, oldCode)
			menuRoles = append(menuRoles, newCode)
			menu.Roles = strings.Join(menuRoles, ",")
			db.Save(&menu)
		}
	}
	users := []model.User{}
	db.Find(&users)
	for _, user := range users {
		userRoles := strings.Split(user.Roles, ",")
		if common.IsContain(userRoles, oldCode) {
			userRoles = common.RemoveStr(userRoles, oldCode)
			userRoles = append(userRoles, newCode)
			user.Roles = strings.Join(userRoles, ",")
			db.Save(&user)
		}
	}
}

// 删除菜单表和用户表对应角色字段
func deleteRoleMenuAndUser(code string) {
	db := model.Db
	menus := []model.Menu{}
	db.Find(&menus)
	for _, menu := range menus {
		menuRoles := strings.Split(menu.Roles, ",")
		if common.IsContain(menuRoles, code) {
			menuRoles = common.RemoveStr(menuRoles, code)
			menu.Roles = strings.Join(menuRoles, ",")
			db.Save(&menu)
		}
	}
	users := []model.User{}
	db.Find(&users)
	for _, user := range users {
		userRoles := strings.Split(user.Roles, ",")
		if common.IsContain(userRoles, code) {
			userRoles = common.RemoveStr(userRoles, code)
			user.Roles = strings.Join(userRoles, ",")
			db.Save(&user)
		}
	}
}

// 注册用户路由
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/get-user-list", getUserList)
	router.POST("/user", addUser)
	router.PUT("/user", updateUser)
	router.DELETE("/user", deleteUser)
	router.POST("/upload-avatar", UploadUserAvatar)
	router.PUT("/user-status",updateUserStatus)
	router.PUT("/user-password",updateUserPassword)
	router.PUT("/user-role",updateUserRoles)
	router.DELETE("/batch-user",deleteBatchUser)
}

//批量删除用户信息
func deleteBatchUser(c *gin.Context) {
	var ids map[string]interface{}
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	idsArr := ids["ids"].([]interface{})
	db := model.Db
	for _, id := range idsArr {
		var user model.User
		db.Where("id = ?", id).First(&user)
		db.Delete(&user)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete user"})
}

//更新用户角色
func updateUserRoles(c *gin.Context) {
	var roleData map[string]interface{}
	err := c.ShouldBindJSON(&roleData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	userId := uint64(roleData["userId"].(float64))
	roles := roleData["roleIds"].([]interface{})
	var roleStr []string
	db := model.Db
	for _, roleId := range roles {
		var role model.Role
		res := db.Where("id = ?", roleId).First(&role)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "role not found"})
			continue
		}
		roleStr = append(roleStr, role.Code)
	}
	roleStrs := strings.Join(roleStr, ",")

	// Implement logic to update a department
	if err := db.Model(&model.User{Id: userId}).Update("roles",roleStrs).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update user failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update user"})
}
type userPassword struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
	Password 			string `gorm:"column:password;type:varchar(255);not null" json:"password"`      	// 角色密码
}
// 表名
func (userPassword) TableName() string {
	return "system_user"
}

//更新用户密码
func updateUserPassword(c *gin.Context) {
	var user userPassword
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	if db.Save(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update user failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update user"})
}


type userStatus struct{
	Id         			uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      	// 角色id
	Status 				uint8  `gorm:"column:status;type:tinyint" json:"status"`      				// 角色状态
}
// 表名
func (userStatus) TableName() string {
	return "system_user"
}

//更新用户状态
func updateUserStatus(c *gin.Context) {
	var user userStatus
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	if db.Debug().Save(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update user failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update user"})
}

//获取用户信息的条件数据
type userCondition struct{
	Username 	string 	`json:"username"`  		// 用户名
	Phone 		string 	`json:"phone"`  		// 电话
	Status 		string 	`json:"status"`  		// 状态
	DeptId 		string 	`json:"deptId"`  		// 部门id
	CurrentPage int 	`json:"currentPage"`  	// 当前页数
	PageSize 	int 	`json:"pageSize"`  		// 每页显示条目个数
	Total 		int 	`json:"total"`  		// 总条目数
	Background 	bool 	`json:"background"`  	// 是否后台查询
}

//获取用户信息
func getUserList(c *gin.Context) {
	var userCondition userCondition
	err := c.ShouldBindJSON(&userCondition)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	users := []model.User{}
	db := model.Db
	if userCondition.Username != "" {
		db = db.Where("username like ?", "%" + userCondition.Username + "%")
	}
	if userCondition.Phone != "" {
		db = db.Where("phone like ?", "%" + userCondition.Phone+ "%")
	}
	if userCondition.Status != "" {
		db = db.Where("status = ?", userCondition.Status)
	}
	if userCondition.DeptId != "" {
		db = db.Where("deptId = ?", userCondition.DeptId)
	}
	var total int64
	db.Model(&model.User{}).Count(&total)
	offset := (userCondition.CurrentPage - 1) * userCondition.PageSize
	db.Limit(userCondition.PageSize).Offset(offset).Find(&users)

	// Implement logic to get departments
	c.JSON(http.StatusOK,
		gin.H{
			"success": true,
			"data": gin.H{
				"list": users,
				"total": total, // 总条目数
		  		"pageSize": userCondition.PageSize, // 每页显示条目个数
		  		"currentPage": userCondition.CurrentPage, // 当前页数
			},
		})
}

//添加用户信息
func addUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	db := model.Db
	if db.Create(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "create user failed"})
		return
	}
	// Implement logic to create a department
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "create user"})
}

//更新用户信息
func updateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to update a department
	if db.Save(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "update user failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update user"})
}

//删除用户信息
func deleteUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	db := model.Db
	// Implement logic to delete a department
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "delete user"})
}

//上传用户头像
func UploadUserAvatar(c *gin.Context) {
	data := make(map[string]interface{})
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "invalid input"})
		return
	}
	id := data["id"].(float64)
	avatar := data["avatar"].(map[string]interface{})
	base64Str := avatar["base64"].(string)
	db := model.Db
	var user model.User
	res := db.Where("id = ?", id).First(&user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "user not found"})
		return
	}
	// Implement logic to update a department
	file,_:= common.SaveBase64Img(base64Str, "avatar")
	user.Avatar = file
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "update user"})
}