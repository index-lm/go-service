package api

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/pkg/model/res"
	"go-service/internal/pkg/sys"
	"go-service/pkg/discovery"
	"go-service/pkg/pb/transfer"
)

func  CreatAccount(c *gin.Context){
	dial := discovery.GetGrpcServiceDial(discovery.ServiceImGateway, "DEFAULT_GROUP", "develop")
	transferClient := transfer.NewTransferClient(dial)
	login, _ := transferClient.PasswordLogin(sys.GetCxt(), &transfer.LoginReq{Username: "123", Password: "pass"})
	res.OkWithData("登录成功！"+login.Success, c)
}
