package service

import (
	"Gin-Template/configs"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf *viper.Viper
var lg *zap.Logger

func InitService() {
	Conf = configs.Conf
	lg = zap.L()
}
