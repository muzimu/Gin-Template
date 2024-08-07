package service

import "Gin-Template/configs"

var Conf *configs.Conf

func InitService(conf *configs.Conf) {
	Conf = conf
}
