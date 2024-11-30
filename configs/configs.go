package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

// 载入配置
func LoadConfig(path string) (*viper.Viper, error) {
	Config = viper.New()
	Config.SetConfigFile(path)
	err := Config.ReadInConfig()
	if err != nil {
		fmt.Println("error occur!", err)
		return Config, err
	}
	return Config, nil
}

func GetConfig() *viper.Viper { return Config }
