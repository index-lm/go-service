package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/pkg/middleware"
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
		authRouter.GET("/login", CreatAccount)
	}
}
