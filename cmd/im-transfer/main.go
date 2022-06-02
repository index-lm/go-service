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
	defer func() {
		i := recover()
		fmt.Println(i)
	}()
	yamlBytes := []byte(im_transfer.YamlStr)
	//解析yaml
	err := yaml.YamlParse(&yamlBytes, &im_transfer.AppConfig)
	if err != nil {
		panic(err.Error())
	}
	// 先从yaml中获取端口
	portStr := fmt.Sprintf("%d", im_transfer.AppConfig.System.Port)
	// 再从启动命令中获取端口参数
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)
	// 初始化系统公共配置
	sys.Initialize(portInt, im_transfer.AppConfig.System.Name)
	// 初始化日志框架
	log.ConfigInit(log.WithLogPath(im_transfer.AppConfig.Log.File),
		log.WithServiceName(sys.ServerName),
		log.WithLogLevel(im_transfer.AppConfig.Log.Level))
	// 初始化orm
	err = db.InitGorm(im_transfer.AppConfig.Mysql.Username,
		im_transfer.AppConfig.Mysql.Password,
		im_transfer.AppConfig.Mysql.Host,
		im_transfer.AppConfig.Mysql.Db,
		im_transfer.AppConfig.Mysql.Conn.MaxIdle,
		im_transfer.AppConfig.Mysql.Conn.MaxIdle,
		model.InitDb)
	if err != nil {
		log.Error("sys", err.Error())
	}
	// 初始化Redis
	db.InitRedis(im_transfer.AppConfig.Redis.Host,
		im_transfer.AppConfig.Redis.Port,
		im_transfer.AppConfig.Redis.Password,
		im_transfer.AppConfig.Redis.Db)
	// 初始化注册中心配置 -注册生产者
	discovery.Initialize(im_transfer.AppConfig.Nacos.IpAddr,
		im_transfer.AppConfig.Nacos.Port,
		im_transfer.AppConfig.Nacos.NamespaceId,
		sys.ServerName,
		im_transfer.AppConfig.Log.File,
		im_transfer.AppConfig.Log.File)
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
