package im_gateway

import _ "embed"

//go:embed application.yaml
var YamlStr string

var AppConfig *Server

type Server struct {
	System System `json:"system" yaml:"system"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Log    Log    `json:"log" yaml:"log"`
	Nacos  Nacos  `json:"nacos" yaml:"nacos"`
}

// System 系统配置
type System struct {
	Port int    `json:"port" yaml:"port"`
	Name string `json:"name" yaml:"name"`
}

// Redis 缓存配置
type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Db       int    `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"poolSize" yaml:"poolSize"`
	Cache    struct {
		TokenExpired int `json:"tokenExpired" yaml:"tokenExpired"`
	}
}

// Log 日志配置
type Log struct {
	File  string `json:"file" yaml:"file"`
	Level string `json:"level" yaml:"level"`
}

// nacos注册中心配置
type Nacos struct {
	IpAddr      string `josn:"ipAddr" yaml:"ipAddr"`
	Port        uint64 `josn:"port" yaml:"port"`
	NamespaceId string `josn:"namespaceId" yaml:"namespaceId"`
}
