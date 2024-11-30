package service

import (
	"Gin-Template/configs"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

func InitService() {
	Conf = configs.GetConfig()
}
