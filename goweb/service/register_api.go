package service

import (
	"fmt"
	"main/config"
	"main/logger"
	"main/middleware"
	"main/module"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "main/docs"

)

type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	err := r.SetTrustedProxies(nil)
	if err != nil {
		logger.Fatalf("Gin set trusted proxies failed! err: #%v", err)
	}
	r.Use(middleware.GinWebLog())
	r.Use(gin.Recovery())
	swagHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/*any", swagHandler)

	authMiddleware := middleware.AuthMiddleWare()

	r.POST("/login", module.LoginHandler)
	r.POST("/refresh-token",module.RefreshTokenHandler)
	r.GET("/get-async-routes", module.GetAsyncRoutes)

	// 部门管理
	module.RegisterDepartmentRoutes(r)
	//菜单管理
	module.RegisterMenuRoutes(r)
	//角色管理
	module.RegisterRoleRoutes(r)
	//用户管理
	module.RegisterUserRoutes(r)
	//日志管理
	module.RegisterLogRoutes(r)
	//产品管理
	module.RegisterProductRoutes(r)
	//供应商管理
	module.RegisterSupplierRoutes(r)
	//代理商管理
	module.RegisterAgentRoutes(r)
	//订单管理
	module.RegisterOrderRoutes(r)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	for _, opt := range options {
		opt(r)
	}
	return r
}

func StartApi() {
	// 初始化路由
	r := Init()
	configBase, err := config.InitConfig()
	if err != nil {
		logger.Fatalf("Get config failed! err: #%v", err)
	}
	fmt.Printf("Listening and serving HTTP on %s\n", configBase.Webapi.Uri)
	if err := r.Run(configBase.Webapi.Uri); err != nil {
		logger.Fatalf("Run web server failed! err: #%v", err)
	}
}
