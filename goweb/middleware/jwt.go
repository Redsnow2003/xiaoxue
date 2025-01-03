package middleware

import (
	"errors"
	"crypto/md5"  // 从 crypto/md5 中导入 Sum() 函数
	"fmt"         // 从 fmt 中导入 Printf() 函数
	"main/logger" // 从 main\logger 中导入 Errorf() 函数
	"main/model"
	"net/http"         // 从 net/http 中导入 StatusOK() 函数
	"strings"
	"time" // 从 time 中导入 Now() 函数
	"gorm.io/gorm"
	jwt "github.com/appleboy/gin-jwt/v2" // 从 github.com/appleboy/gin-jwt/v2 中导入 GinJWTMiddleware() 函数
	"github.com/gin-gonic/gin"           // 从 github.com/gin-gonic/gin 中导入 Gin() 函数
)

// 用户信息
type JwtUser struct {
	UserName string
}

// 身份验证的 key 值
var identityKey = "id"

// 登录信息
type login struct {
    UserName 	string `form:"username" json:"username" binding:"required"`
	Password 	string `form:"password" json:"password" binding:"required"`
}

// MD5加密
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// 登录
func AuthMiddleWare() *jwt.GinJWTMiddleware {
	var userInfo *model.User
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		// 中间件名称
		Realm: "gin-jwt",
		Key: []byte("secret key"),
		// token 过期时间
		Timeout: 24 * time.Hour,
		// token 刷新最大时间
		MaxRefresh: 24 * time.Hour,
		// 身份验证的 key 值
		IdentityKey: identityKey,
		// 登录期间的回调的函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(JwtUser); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		// 解析并设置用户身份信息
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if claims[identityKey] == nil {
				return nil
			}
			return JwtUser{
				UserName: claims[identityKey].(string),
			}
		},
		// 根据登录信息对用户进行身份验证的回调函数
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVars login
			if err := c.ShouldBind(&loginVars); err != nil {
				logger.Errorf("登录信息错误：%v",err)
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVars.UserName
			password := loginVars.Password
			logger.Infof("username:%v,password:%v",username,password)
			res := SelelctByUserName(username)
			userInfo = res
			if res != nil && (password) == (res.Password) {		
				return JwtUser{
					UserName: username,
				},nil
			}
			return nil,jwt.ErrFailedAuthentication
		},
		// 接收用户信息并编写授权规则
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(JwtUser); ok {
				return true
			}
			return false
		},
		// 自定义未授权的回调函数
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code,gin.H{
				"code":code,
				"success":  false,
				"message":message,
				"data":gin.H {},
			})
		},
		// 自定义登录成功的回调函数
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			// 生成刷新 token
			roles := strings.Split(userInfo.Roles, ",")
			permissions := strings.Split(userInfo.Permissions, ",")
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"success":  true,
				"message": "登录成功",
				"data": gin.H {
					"avatar"		: 	userInfo.Avatar,
					"id"			:	userInfo.Id,
					"username"		:	userInfo.UserName,
					"nickname"		: 	userInfo.Nickname,
					"roles"			: 	roles,
					"permissions"	:	permissions,
					"accessToken"	:	token,
					"refreshToken"	:	token,
					"expire"		: 	expire.Format(time.RFC3339),
				},
			})
		},
		// token 刷新成功的回调函数
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"success":  true,
				"message": "刷新成功",
				"data": gin.H {
					"accessToken"	:	token,
					"refreshToken"	:	token,
					"expire"		: 	expire.Format(time.RFC3339),
				},
			})
		},
		// token 获取函数
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	if err != nil {
		logger.Errorf("jwt 中间件创建错误：%v",err)
	}
	// 登录成功的回调函数
	return authMiddleWare
}

// 根据用户名查询用户
func SelelctByUserName(username string) *model.User {
	db := model.Db
	u := model.User{}
	res := db.Where("username = ?", username).First(&u)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logger.Debugf("Select by username err:" + "未查找到相关数据")
		return nil
	}
	return &u
}

// 获取当前用户
func GetCurrentUser(c* gin.Context) *model.User{

	// 解析出请求中的token
	token,err := AuthMiddleWare().ParseToken(c)
	if(err != nil){
		logger.Debugf("Get current user err:%v", err)
		return nil
	}
	// 从token中解析出用户信息
	claims := jwt.ExtractClaimsFromToken(token)
	// 从用户信息中解析出用户账号
	username := claims["id"].(string)
	// 根据用户账号查询用户信息
	return SelelctByUserName(username)
}