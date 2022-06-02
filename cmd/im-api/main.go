package main

import (
	"flag"
	"fmt"
	im_api "go-service/configs/im-api"
	"go-service/internal/app/im-api/config"
	"go-service/internal/pkg/sys"
	"go-service/pkg/db"
	"go-service/pkg/discovery"
	"go-service/pkg/log"
	"go-service/pkg/yaml"
	"strconv"
)

func main() {
	defer func() {
		i := recover()
		fmt.Println(i)
	}()
	yamlBytes := []byte(im_api.YamlStr)
	err := yaml.YamlParse(&yamlBytes, &im_api.AppConfig)
	if err != nil {
		panic(err.Error())
	}
	// 先从yaml中获取端口
	portStr := fmt.Sprintf("%d", im_api.AppConfig.System.Port)
	// 再从启动命令中获取端口参数
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)
	// 初始化系统公共配置
	sys.Initialize(portInt, im_api.AppConfig.System.Name)
	// 初始化日志框架
	log.ConfigInit(log.WithLogPath(im_api.AppConfig.Log.File),
		log.WithServiceName(sys.ServerName),
		log.WithLogLevel(im_api.AppConfig.Log.Level))
	// 初始化Redis
	db.InitRedis(im_api.AppConfig.Redis.Host, im_api.AppConfig.Redis.Port, im_api.AppConfig.Redis.Password, im_api.AppConfig.Redis.Db)
	// 初始化注册中心配置 -注册生产者
	discovery.Initialize(im_api.AppConfig.Nacos.IpAddr,
		im_api.AppConfig.Nacos.Port,
		im_api.AppConfig.Nacos.NamespaceId,
		sys.ServerName,
		im_api.AppConfig.Log.File,
		im_api.AppConfig.Log.File)

	err = config.InitWeb(sys.ServerPort)
	if err != nil {
		log.Error("sys", err.Error())
	}
}

//go func() {
//	//创建监听退出chan
//	c := make(chan os.Signal)
//	//监听指定信号 ctrl+c kill
//	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//	s := <-c
//	fmt.Println("程序推出--------------------", s)
//	fmt.Println("Start Exit...")
//	fmt.Println("Execute Clean...")
//	fmt.Println("End Exit...")
//	os.Exit(0)
//}()
