package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type InitRouter func(router *gin.Engine)

func InitGin(port uint64,initRouter InitRouter)  {
	address := fmt.Sprintf(":%d", port)
	// 默认已经连接了 Logger and Recovery 中间件
	var router = gin.New()
	// 初始化路由
	initRouter(router)
	err := router.Run(address)
	if err != nil {
		panic(err)
	}
}
