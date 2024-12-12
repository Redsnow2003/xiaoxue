package module

import (
	"github.com/gin-gonic/gin"
	"main/middleware"
)
//@Description 用户登录
//@Summary 获取账号进行登录
//@Accept multipart/form-data
//@Produce application/json
//@Param username formData string true "用户名"
//@Param password formData string true "密码"
//@Success 200 {json} json "{"code": 200,"data": {"id": 2,"img": "","message": "success", "name": "管理员", "role": 1,"username": "admin" }, "expire": "expire-time", "token": "token"}"
//@Failure 500 "获取账号信息出错"
//@Failure 404 "未找到此用户"
//@Router /login [POST]
func LoginHandler(c *gin.Context) {
	middleware.AuthMiddleWare().LoginHandler(c)
}
