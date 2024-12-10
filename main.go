package main

import (
	"Gin-Template/configs"
	"Gin-Template/logger"
	"Gin-Template/router"
	"Gin-Template/service"
	"log"

	"github.com/spf13/cobra"
)

var configPath string

func init() {
	var rootCmd = &cobra.Command{
		Use: "app",
	}
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "configs/dev.yaml", "config file path")
	rootCmd.Execute()
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

	// 初始化日志配置 使用中间件
	logger.InitLogger(conf)
	service.InitService()
	router.StartServer()
}
