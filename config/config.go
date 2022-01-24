package config

import "time"

// 配置
type Config struct {
	Database Database
	Redis    Redis
	System   System
	Server   Server
}

// 系统配置信息
type System struct {
	PageSize uint32
}

type Server struct {
	BaseUrl      string
	HttpPort     uint32
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	CloseTimeout time.Duration
}

// 关系型数据库配置信息
type Database struct {
	Type        string
	Host        string
	Port        uint32
	User        string
	Password    string
	Name        string
	TablePrefix string
}

// redis 配置信息
type Redis struct {
	Host        string
	Db          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	MaxRetries  int
}
