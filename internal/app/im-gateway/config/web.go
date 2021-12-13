package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/app/im-gateway/api"
)

func InitWeb(port int) error {
	address := fmt.Sprintf(":%d", port)
	// 默认已经连接了 Logger and Recovery 中间件
	var router = gin.New()
	// 初始化路由
	api.InitRouter(router)
	err := router.Run(address)
	if err != nil {
		return err
	}
	return nil
}
