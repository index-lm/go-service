package main

import (
	_ "embed"
	im_gateway "go-service/configs/im-gateway"
	"go-service/internal/app/im-gateway/config"
	"go-service/pkg/log"
	"go-service/pkg/yaml"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var myConfig *config.Server
	yamlBytes := []byte(im_gateway.ImGatewayYaml)
	err := yaml.YamlParse(&yamlBytes, &myConfig)
	if err != nil {
		log.Error("sys", err.Error())
	}
	log.Init("/opt/go", "info", 200, 30, 90, false, "im-gateway")
	//discovery.Init()
	wg.Wait()
}
