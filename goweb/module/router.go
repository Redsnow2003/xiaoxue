package module

import (
	"main/model"
	"main/common"
	"net/http"
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
func GetAsyncRoutes(c *gin.Context) {
	menu1 := findMenuItems(0)
	if menu1 == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, 
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
			roles_len := len(roles)
			roles_new := make([]string, 0)
			for i := 0; i < roles_len; i++ {
				if roles[i] != "" {
					roles_new = append(roles_new, roles[i])
				}
			}
			if !common.IsContain(roles_new, "admin") {
				roles_new = append(roles_new, "admin")
			}
			
			meta2["roles"] = roles_new;

			auths := strings.Split(item2.Auths, ",")
			auths_len := len(auths)
			auths_new := make([]string, 0)
			for i := 0; i < auths_len; i++ {
				if auths[i] != "" {
					auths_new = append(auths_new, auths[i])
				}
			}
			if(len(auths_new) > 0) {
				meta2["auths"] = auths_new;
			}
			
			d2["meta"] = meta2;
			children = append(children, d2)
		}
		d1["children"] = children
		data = append(data, d1)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, 
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