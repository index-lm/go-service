package middleware

import (
	"github.com/gin-gonic/gin"
	"go-service/internal/pkg/model/res"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//封装通用json返回
			res.FailWithDetailed(0, nil, errorToString(r), c)
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
