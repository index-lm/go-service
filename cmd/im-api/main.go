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
	var myConfig *config.Server
	yamlBytes := []byte(im_api.ImApiYaml)
	err := yaml.YamlParse(&yamlBytes, &myConfig)
	if err != nil {
		log.Error("sys", err.Error())
	}
	portStr := fmt.Sprintf("%d", myConfig.System.Port)
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)
	sys.Init(portInt, myConfig.System.Name)
	log.Initialize("/opt/go", "info", 200, 30, 90, false, sys.ServerName)

	// 初始化Redis
	db.InitRedis(myConfig.Redis.Host, myConfig.Redis.Port, myConfig.Redis.Password, myConfig.Redis.Db)
	discovery.Initialize(sys.ServerPort, sys.ServerName)

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
