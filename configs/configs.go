package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

// 载入配置
func LoadConfig(path string) (*viper.Viper, error) {
	Conf = viper.New()
	Conf.SetConfigFile(path)
	err := Conf.ReadInConfig()
	if err != nil {
		fmt.Println("error occur!", err)
		return Conf, err
	}
	return Conf, nil
}
