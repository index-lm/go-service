package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var GoRedis *redis.Client

func Set() {

}

// InitRedis 初始化缓存（redis）
func InitRedis(host string, port string, password string, db int) {
	sprintf := fmt.Sprintf("%s:%s", host, port)
	client := redis.NewClient(&redis.Options{
		Addr:     sprintf,
		Password: password,
		DB:       db,
	})
	GoRedis = client
}
