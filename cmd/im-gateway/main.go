package main

import (
	_ "embed"
	"fmt"
	im_gateway "go-service/configs/im-gateway"
	"gopkg.in/yaml.v3"
)

func main() {
	yaml1 := im_gateway.ImGatewayYaml
	fmt.Println(yaml1)
	m := make(map[interface{}]interface{})
	_ = yaml.Unmarshal([]byte(yaml1), &m)
	for k, v := range m {
		fmt.Printf("--- k: %v -v: %v \n", k, v)
	}
}
