package db

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"tstj-based-on-AI/config"
)

type DbStruct struct {
	Mysql *gorm.DB
	Redis *redis.Client
}

var FlagConfig = flag.String("f", "./config/config.yaml", "choose config file (.yaml)")
var Db = &DbStruct{}

func InitDb(ctx *gin.Context) {
	serverConfig := config.ServerConfig{}

	config.GetConfigMessageFromYaml(FlagConfig, &serverConfig, ctx)

	Db = &DbStruct{Mysql: InitMysql(ctx), Redis: InitRedis(ctx)}
}
