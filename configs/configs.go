package configs

import (
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
)

// 配置文件
type Conf struct {
	// 版本
	Version string `yaml:"-"`
	// 配置文件目录
	ConfigDir string `yaml:"-"`
	// 运行模式
	Runmode string `yaml:"runmode"`
	// 应用相关配置
	App struct {
		APPID     string `yaml:"APPID"`
		SecretKey string `yaml:"SecretKey"`
		Cache     bool   `yaml:"Cache"`
	} `yaml:"app"`
	// DB *db.Config `yaml:"db"`
	// Redis    *db.RedisConf `yaml:"redis"`
	LogfileEnable  bool              `yaml:"logfile_enable"`
	LogfileSetting lumberjack.Logger `yaml:"logfile_setting"` // 日志文件存放地址,未设置的时候 直接控制台输出
	Http           struct {
		Port           string   `yaml:"port"` // http端口
		AllowedOrigins []string `yaml:"allowed_origins"`
		UrlPrefix      string   `yaml:"url_prefix"` //
	} `yaml:"http"`
}

// 载入配置
func LoadConfig(path string) (*Conf, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var res Conf
	err = yaml.Unmarshal(bs, &res)
	if err != nil {
		return nil, err
	}

	res.ConfigDir = filepath.Dir(path)

	// 读取同目录下的version文件(若没有  版本为空)
	versionPath := filepath.Join(res.ConfigDir, "version")
	bs, _ = os.ReadFile(versionPath)
	res.Version = string(bs)

	return &res, nil
}
