package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-service/internal/pkg/middleware"
	"go-service/internal/pkg/model/res"
	"go-service/pkg/discovery"
	"go-service/pkg/log"
	"go-service/pkg/pb/transfer"
	"google.golang.org/grpc"
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
		authRouter.GET("/login", func(c *gin.Context) {
			instance := discovery.SelectOneHealthyInstance(*discovery.ServerCenter, vo.SelectOneHealthInstanceParam{
				ServiceName: "im-transfer",
				GroupName:   "DEFAULT_GROUP",
				Clusters:    []string{"develop"},
			})
			address := fmt.Sprintf("%s:%d", instance.Ip, instance.Port)
			dial, err := grpc.Dial(address, grpc.WithInsecure())

			if err != nil {
				log.Error("grpc", err.Error())
			}
			transferClient := transfer.NewTransferClient(dial)
			login, err := transferClient.PasswordLogin(context.Background(), &transfer.LoginReq{Username: "123", Password: "pass"})
			res.OkWithData("登录成功！"+login.Success, c)
		})
	}
}
