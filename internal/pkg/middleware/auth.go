package middleware

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/pkg/model/res"
)

func Auth(context *gin.Context) {
	header := context.GetHeader("X-Access-Token")
	if header == "" {
		//封装通用json返回
		res.FailWithDetailed(res.AUTH, nil, "无效的请求头", context)
		//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
		context.Abort()
		return
	}
	//err := utils.JwtVerification(header, context)
	//if err != nil {
	//	//封装通用json返回
	//	res.FailWithDetailed(res.AUTH, nil, err.Error(), context)
	//	//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
	//	context.Abort()
	//	return
	//}
	context.Next()
}
