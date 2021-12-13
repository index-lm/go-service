package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/internal/pkg/middleware"
	"go-service/internal/pkg/model/res"
	"go-service/pkg/log"
)

func InitRouter(router *gin.Engine) {
	//全局中间件
	//全局异常处理
	router.Use(middleware.Recover)
	//跨域请求放行中间件
	router.Use(middleware.Cors())
	//
	//routerGroupSecurity := router.Group("/security")
	////认证中间件
	//routerGroupSecurity.Use(middleware.Auth)
	routerGroupCommon := router.Group("/common")
	//路由注册
	routerCommonInit(routerGroupCommon)
	//routerSecurityInit(routerGroupSecurity)
}

//
func routerCommonInit(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.GET("/login", func(context *gin.Context) {
			go testLog()
			go testLog()
			go testLog()
			go testLog()
			go testLog()
			go testLog()
			go testLog()
			res.OkWithData("登录成功！", context)
		})
	}
}
func testLog() {
	for i := 0; i < 99999; i++ {
		sprintf := fmt.Sprintf("12312%d", i)
		log.Info("tttt", sprintf)
	}
}
