package main

import (
	"Gin-Template/configs"
	"Gin-Template/logger"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "configs/config.yml", "配置文件地址")
	flag.Parse()
}

func main() {
	r := gin.Default()

	// 读取配置
	conf, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Fatalln("读取配置文件失败：", err.Error())
		return
	}
	fmt.Println(conf.App.APPID)

	// 初始化日志配置 使用中间件
	logger.InitLogger()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	zap.L().Info("Starting server...")

	r.Run(":80")
}
