package main

import (
	_ "embed"
	im_gateway "go-service/configs/im-gateway"
	"go-service/internal/app/im-gateway/config"
	"go-service/pkg/db"
	"go-service/pkg/log"
	"go-service/pkg/yaml"
)

func main() {
	var myConfig *config.Server
	yamlBytes := []byte(im_gateway.ImGatewayYaml)
	err := yaml.YamlParse(&yamlBytes, &myConfig)
	if err != nil {
		log.Error("sys", err.Error())
	}
	log.InitLogger("/opt/go", "info", 200, 30, 90, false, "gateway")
	// 初始化orm
	err = db.InitGorm(myConfig.Mysql.Username, myConfig.Mysql.Password, myConfig.Mysql.Host, myConfig.Mysql.Db, myConfig.Mysql.Conn.MaxIdle, myConfig.Mysql.Conn.MaxIdle)
	if err != nil {
		log.Error("sys", err.Error())
	}
	// 初始化Redis
	db.InitRedis(myConfig.Redis.Host, myConfig.Redis.Port, myConfig.Redis.Password, myConfig.Redis.Db)
	err = config.InitWeb(myConfig.System.Port)
	if err != nil {
		log.Error("sys", err.Error())
	}
}
