package main

import (
	_ "embed"
	"flag"
	"fmt"
	im_transfer "go-service/configs/im-transfer"
	"go-service/internal/app/im-transfer/config"
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
	var myConfig *config.Server
	yamlBytes := []byte(im_transfer.ImTransferYaml)
	err := yaml.YamlParse(&yamlBytes, &myConfig)
	if err != nil {
		log.Error("sys", err.Error())
	}
	portStr := fmt.Sprintf("%d", myConfig.System.Port)
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)

	sys.Init(portInt, myConfig.System.Name)
	log.Init("/opt/go", "info", 200, 30, 90, false, sys.ServerName)
	// 初始化orm
	err = db.InitGorm(myConfig.Mysql.Username, myConfig.Mysql.Password, myConfig.Mysql.Host, myConfig.Mysql.Db, myConfig.Mysql.Conn.MaxIdle, myConfig.Mysql.Conn.MaxIdle)
	if err != nil {
		log.Error("sys", err.Error())
	}
	// 初始化Redis
	db.InitRedis(myConfig.Redis.Host, myConfig.Redis.Port, myConfig.Redis.Password, myConfig.Redis.Db)

	discovery.Init(sys.ServerPort, sys.ServerName)
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
