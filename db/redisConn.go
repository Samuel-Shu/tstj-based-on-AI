package db

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"tstj-based-on-AI/config"
)

func InitRedis(c *gin.Context) *redis.Client {
	var config1 any
	var exist bool
	if config1, exist = c.Get("ServerConfig"); !exist {
		log.Fatal(errors.New("redis配置信息加载失败！！"))
	}
	conf := config1.(*config.ServerConfig)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.RdbPassword, // no password set
		DB:       conf.Num,         // use default DB
	})
	return rdb
}
