package main

import (
	_ "embed"
	"flag"
	"fmt"
	im_transfer "go-service/configs/im-transfer"
	"go-service/internal/app/im-transfer/model"
	"go-service/internal/app/im-transfer/service"
	"go-service/internal/pkg/sys"
	"go-service/pkg/db"
	"go-service/pkg/discovery"
	"go-service/pkg/log"
	"go-service/pkg/pb/transfer"
	"go-service/pkg/yaml"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	yamlBytes := []byte(im_transfer.ImTransferYaml)
	//解析yaml
	err := yaml.YamlParse(&yamlBytes, im_transfer.AppConfig)
	if err != nil {
		log.Error("sys", err.Error())
		return
	}
	// 先从yaml中获取端口
	portStr := fmt.Sprintf("%d", im_transfer.AppConfig.System.Port)
	// 再从启动命令中获取端口参数
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)
	// 初始化系统公共配置
	sys.Initialize(portInt, im_transfer.AppConfig.System.Name)
	// 初始化日志框架
	log.Initialize("/opt/go", "info", 200, 30, 90, false, sys.ServerName)
	// 初始化orm
	err = db.InitGorm(im_transfer.AppConfig.Mysql.Username, im_transfer.AppConfig.Mysql.Password, im_transfer.AppConfig.Mysql.Host, im_transfer.AppConfig.Mysql.Db, im_transfer.AppConfig.Mysql.Conn.MaxIdle, im_transfer.AppConfig.Mysql.Conn.MaxIdle, model.InitDb)
	if err != nil {
		log.Error("sys", err.Error())
	}
	// 初始化Redis
	db.InitRedis(im_transfer.AppConfig.Redis.Host, im_transfer.AppConfig.Redis.Port, im_transfer.AppConfig.Redis.Password, im_transfer.AppConfig.Redis.Db)

	discovery.Initialize(sys.ServerPort, sys.ServerName)
	server := grpc.NewServer()
	transfer.RegisterTransferServer(server, new(service.Transfer))
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", sys.ServerPort))
	if err != nil {
		log.Error("服务监听端口失败", err.Error())
	}
	//reflection.Register(server)
	err = server.Serve(listen)
	if err != nil {
		log.Error("服务监听端口失败", err.Error())
	}
}
