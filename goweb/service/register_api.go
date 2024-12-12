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
	r.GET("/system/getNodes", module.GetNodes)
	r.GET("/system/getNodeInfo", module.GetNodeInfo)
	r.POST("/system/test", module.Test)
	r.POST("/system/controlNode", module.ControlNode)
	r.POST("/system/controlProcess", module.ControlProcess)
	r.GET("/system/getProcesses", module.GetProcesses)
	r.GET("/device/getDevices", module.GetDevices)
	r.GET("/device/getDeviceChannels", module.GetDeviceChannels)
	r.GET("/device/getAnalogPoints", module.GetAnalogPoints)
	r.GET("/device/getDigitalPoints", module.GetDigitalPoints)
	r.GET("/device/getProtocolParam", module.GetProtocolParam)
	r.POST("/device/startSimulator", module.StartSimulator)
	r.GET("/device/stopSimulator", module.StopSimulator)


	r.GET("/ws", module.WsHandler)
	r.GET("/ws/wsGetRealTimeData", module.WsGetRealTimeData)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	Include(userRouter)

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
