package main

import (
	"Gin-Template/configs"
	"Gin-Template/logger"
	"Gin-Template/router"
	"Gin-Template/service"
	"flag"
	"fmt"
	"log"

	"go.uber.org/zap"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "configs/config.yml", "配置文件地址")
	flag.Parse()
}

// @title                       Gin-Template API接口文档
// @version                     v1.0
// @description                 Gin-Template后端API	  作者：木子木先生
func main() {

	// 读取配置
	conf, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Fatalln("读取配置文件失败：", err.Error())
		return
	}
	fmt.Println(conf.App.APPID)

	// 初始化日志配置 使用中间件
	logger.InitLogger()
	// r.Use(logger.GinLogger(), logger.GinRecovery(true))
	zap.L().Info("Starting server...")
	service.InitService(conf)
	router.StartServer(conf)
}
