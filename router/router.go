package router

import (
	"Gin-Template/configs"
	_ "Gin-Template/docs"
	"Gin-Template/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
)

var Conf *viper.Viper
var lg *zap.Logger

func StartServer() {
	Conf = configs.Conf
	lg = zap.L()
	r := gin.Default()
	r.Use(func() gin.HandlerFunc {
		config := cors.DefaultConfig()
		config.AllowHeaders = []string{"*"}
		config.AllowOrigins = Conf.GetStringSlice("http.allowed_origins")
		config.AllowCredentials = true
		return cors.New(config)
	}())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := Conf.GetString("http.port")
	if port != "" {
		lg.Info("Server start", zap.String("port", port))
		r.Run(port)
	} else {
		lg.Info("Server start on default port", zap.String("port", ":8080"))
		r.Run(":8080")
	}
}
