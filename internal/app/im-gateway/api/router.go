package api

import (
	"github.com/gin-gonic/gin"
	"service-im/handler"
	"service-im/utils"
)

func routerCommonInit(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("/auth")
	{
		AuthRouter.POST("/login", handler.Login)
		AuthRouter.GET("/login/mobile/code", handler.GetMobileLoginCode)
		AuthRouter.POST("/login/mobile/code", handler.MobileLoginCode)
	}
	AreaRouter := Router.Group("/area")
	{
		AreaRouter.GET("/district/list", handler.DistrictList)
		AreaRouter.GET("/branch/list", handler.BranchList)
		AreaRouter.POST("/branch/import", handler.BranchImport)
		AreaRouter.GET("/branch/adminReport", handler.BranchAdminReport)
		AreaRouter.POST("/doctor/import", handler.DoctorImport)
		AreaRouter.POST("/test", func(context *gin.Context) {
			//
			err := utils.MqSend("test", "a", "hello", 0)
			if err != nil {
				panic(err.Error())
			}
		})
	}
}
func routerSecurityInit(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("/auth")
	{
		BaseRouter.DELETE("/logout", handler.Logout)
	}
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.GET("/list", handler.OrderList)
		OrderRouter.GET("/details", handler.OrderDetails)
		OrderRouter.POST("/modifyExt", handler.ModifyOrderExt)
		OrderRouter.POST("/createRoom", handler.CreateRoom)
	}
	DoctorList := Router.Group("/doctor")
	{
		DoctorList.GET("/list", handler.GetDoctorList)
		DoctorList.GET("/details", handler.OrderDetails)
	}
	PatientRouter := Router.Group("/patient")
	{
		PatientRouter.GET("/list", handler.PatientList)
	}
}
