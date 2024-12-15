package module

import (
	"main/model"
	"github.com/gin-gonic/gin"
	"strings"
)

//@Description 获取异步路由
//@Summary 获取异步路由
//@Accept multipart/form-data
//@Produce application/json
//@Success 200 {json} json "{"code": 200,"data": [{"id": 1,"name": "首页","path": "/home","icon": "el-icon-s-home","children": []}]}"
//@Failure 500 "获取异步路由出错"
//@Router /get-async-routes [GET]

/*
const permissionRouter = {
  path: "/permission",
  meta: {
    title: "menus.purePermission",
    icon: "ep:lollipop",
    rank: permission
  },
  children: [
    {
      path: "/permission/page/index",
      name: "PermissionPage",
      meta: {
        title: "menus.purePermissionPage",
        roles: ["admin", "common"]
		c.JSON(200, gin.H{
			"code": 200,
			"success": true,
			"data": data,
		})
	}
    },
    {
      path: "/permission/button",
      meta: {
        title: "menus.purePermissionButton",
        roles: ["admin", "common"]
      },
      children: [
        {
          path: "/permission/button/router",
          component: "permission/button/index",
          name: "PermissionButtonRouter",
          meta: {
            title: "menus.purePermissionButtonRouter",
            auths: [
              "permission:btn:add",
              "permission:btn:edit",
              "permission:btn:delete"
            ]
          }
        },
        {
          path: "/permission/button/login",
          component: "permission/button/perms",
          name: "PermissionButtonLogin",
          meta: {
            title: "menus.purePermissionButtonLogin"
          }
        }
      ]
    }
  ]
};
*/
func GetAsyncRoutes(c *gin.Context) {
	menu1 := findMenuItems(0)
	if menu1 == nil {
		c.JSON(500, gin.H{
			"code": 500, 
			"success": true,
			"data": []model.Menu{},
		})
		return
	}
	data := make([]interface{}, 0)
	for _, item := range menu1 {
		d1 := make(map[string]interface{})
		d1["path"] = item.Path;
		meta := make(map[string]interface{})
		meta["title"] = item.Title;
		meta["icon"] = item.Icon;
		meta["rank"] = item.Rank;
		d1["meta"] = meta;
		menu2 := findMenuItems(item.Id)
		children := make([]interface{}, 0)
		for _, item2 := range menu2 {
			d2 := make(map[string]interface{})
			d2["path"] = item2.Path;
			d2["name"] = item2.Name;
			d2["component"] = item2.Component;
			meta2 := make(map[string]interface{})
			meta2["title"] = item2.Title;
			meta2["icon"] = item2.Icon;
			roles := strings.Split(item2.Roles, ",")
			
			if len(roles) == 0 {
				roles = append(roles, "admin")
			}
			
			meta2["roles"] = roles;
			auths := strings.Split(item2.Auths, ",")
			if(auths != nil) {
				meta2["auths"] = auths;
			}
			
			d2["meta"] = meta2;
			children = append(children, d2)
		}
		d1["children"] = children
		data = append(data, d1)
	}
	c.JSON(200, gin.H{
		"code": 200, 
		"success": true,
		"data": data,
	})
}
func findMenuItems(id uint64) []model.Menu {
	db := model.Db
	menu := []model.Menu{}
	db.Where("parentId = ?", id).Find(&menu)
	return menu
}