package db

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"tstj-based-on-AI/config"
)

func InitMysql(c *gin.Context) *gorm.DB {
	var config1 any
	var exist bool
	if config1, exist = c.Get("ServerConfig"); !exist {
		log.Fatal(errors.New("mysql配置信息加载失败！！"))
	}
	conf := config1.(*config.ServerConfig)
	dsn := conf.Name + ":" + conf.Password + "@tcp(" + conf.Ip + ":" + conf.Port + ")/" + conf.DatabaseName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("mysql connect failed!", err)
	}
	return db
}
