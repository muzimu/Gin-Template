package router

import (
	"Gin-Template/configs"
	_ "Gin-Template/docs"
	"Gin-Template/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var Conf *configs.Conf

func StartServer(conf *configs.Conf) {
	Conf = conf
	r := gin.Default()
	r.Use(func() gin.HandlerFunc {
		config := cors.DefaultConfig()
		// config.AllowHeaders = append(config.AllowHeaders, conf.Http.AllowedHeaders...)
		config.AllowHeaders = []string{"*"}
		config.AllowOrigins = Conf.Http.AllowedOrigins
		config.AllowCredentials = true
		return cors.New(config)
	}())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":80")
}
