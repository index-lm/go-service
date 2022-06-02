package main

import (
	_ "embed"
	"flag"
	"fmt"
	im_gateway "go-service/configs/im-gateway"
	"go-service/internal/pkg/sys"
	"go-service/pkg/yaml"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	defer func() {
		i := recover()
		fmt.Println(i)
	}()
	yamlBytes := []byte(im_gateway.YamlStr)
	//解析yaml
	yaml.YamlParse(&yamlBytes, &im_gateway.AppConfig)
	// 先从yaml中获取端口
	portStr := fmt.Sprintf("%d", im_gateway.AppConfig.System.Port)
	// 再从启动命令中获取端口参数
	portInt, _ := strconv.ParseUint(*flag.String("port", portStr, "启动端口"), 10, 64)
	// 初始化系统公共配置
	sys.Initialize(portInt, im_gateway.AppConfig.System.Name)
	//discovery.Initialize()
	wg.Wait()
}
