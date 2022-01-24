package main

import "notification/config"

var Conf = &config.Config{}

func main() {
	// 加载配置文件
	configLoader := config.NewLoader()
	configLoader.InitConfig()
}
