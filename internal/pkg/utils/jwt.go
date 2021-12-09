package utils

import (
	"github.com/gin-gonic/gin"
	"service-im/config"
	"service-im/global"
	"time"
)

//生成令牌
func JwtGenerate(userId uint64) (string, error) {
	claims := config.CustomClaims{
		UserId: userId,
	}
	unix := time.Now().Unix()
	// 签名生效时间
	claims.NotBefore = unix - 1000
	// 过期时间 一小时
	claims.ExpiresAt = unix
	//签发人
	claims.Issuer = "indexLm"
	jwt, err := global.MyJwt.Create(claims)
	if err != nil {
		return "生成令牌失败", err
	}
	return jwt, nil
}

func JwtVerification(tokenString string, c *gin.Context) error {
	parse, err := global.MyJwt.Parse(tokenString)
	if err != nil {
		return err
	}
	if c.Keys == nil {
		c.Keys = make(map[string]interface{}, 0)
	}
	//unix := time.Now().Unix()
	//if (unix - parse.ExpiresAt) > 0 {
	//	return config.TokenExpired
	//}
	c.Keys["userId"] = parse.UserId
	c.Keys["branchId"] = parse.BranchId
	return nil
}
