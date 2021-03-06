package config

import (
	"os"
	"path"

	"github.com/lujiang618/gtools/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf = &Config{}

func getConfigFileName() string {

	return path.Join(utils.GetRootPath("notification"), "config.yaml")
}

type Loader struct {
	ConfigFile string
	Parser     *viper.Viper
}

func NewLoader() *Loader {
	return &Loader{
		ConfigFile: getConfigFileName(),
		Parser:     viper.New(),
	}
}

// 初始化配置文件
func (l *Loader) InitConfig() {

	// 检查文件是否存在
	if exist := utils.IsValidPath(l.ConfigFile); !exist {
		logrus.WithField("config-path", l.ConfigFile).Error("找不到配置文件")
		os.Exit(0)
	}

	// 读取配置文件
	l.Parser.SetConfigType("yaml")
	l.Parser.AddConfigPath(".")
	l.Parser.SetConfigFile(l.ConfigFile)

	l.LoadConfig()
}

// 加载配置文件
func (l *Loader) LoadConfig() {
	err := l.Parser.ReadInConfig()
	if err != nil {
		logrus.WithField("config-path", l.ConfigFile).Error(err)
		os.Exit(0)
	}

	err = l.Parser.Unmarshal(Conf)
	if err != nil {
		logrus.WithField("config-path", l.ConfigFile).Error(err)
	} else {
		logrus.WithField("config-path", l.ConfigFile).Info("成功读取配置文件")
	}
}
