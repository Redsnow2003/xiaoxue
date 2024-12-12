package main

import (
	"fmt"
	"main/config"
	"main/logger"
	"main/service"
)

var configInfo *config.Config

// @title Swagger Example API
// @version 1.0
// @description this is a sample server celler server
// @termsOfService https://www.swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 17732235526@163.com
 
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host 127.0.0.1:8080
// @BasePath /api
func main() {
	RedisSubService := &service.RedisSubService{}
	RedisSubService.StartRedisSubService()
	service.StartApi()
}

func init() {
	// 初始化配置文件
	var err error
	configInfo, err = config.InitConfig()
	if err != nil {
		fmt.Printf("Init config err %v", err)
	}
	// 为日志配置文件赋值
	configInit := logger.LogConfig{
		EnableConsole:     configInfo.Log.EnableConsole,
		ConsoleJSONFormat: configInfo.Log.ConsoleJSONFormat,
		ConsoleLevel:      configInfo.Log.ConsoleLevel,
		EnableFile:        configInfo.Log.EnableFile,
		FileJSONFormat:    configInfo.Log.FileJSONFormat,
		FileLevel:         configInfo.Log.FileLevel,
		FileLocation:      configInfo.Log.FileLocation,
		MaxAge:            configInfo.Log.MaxAge,
		MaxSize:           configInfo.Log.MaxSize,
		Compress:          configInfo.Log.Compress,
	}
	// 初始化日志
	err = logger.InitGlobalLogger(configInit)
	if err != nil {
		fmt.Printf("Init logger err %v", err)
	}
}
