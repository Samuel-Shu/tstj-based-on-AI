package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tstj-based-on-AI/api"
	"tstj-based-on-AI/db"
	"tstj-based-on-AI/middleware"
	"tstj-based-on-AI/model"
)

func InitRouter() {
	r := gin.Default()
	// 允许跨域请求
	r.Use(middleware.Cors())
	//定义无需鉴权路由组,并初始化配置信息，写进*gin.context里
	NoAuthAPI := r.Group("/api", db.InitDb)
	//定义需鉴权路由组
	AuthAPI := r.Group("/api", db.InitDb)
	AuthAPI.Use(middleware.JWT())

	/*
		定义一个测试路由  /demo/
	*/
	NoAuthAPI.POST("/demo/", func(c *gin.Context) {
		// 初次启动程序时，若无sql文件，可执行该接口以初始化数据表
		if err := db.Db.Mysql.AutoMigrate(&model.User{}); err != nil {
			log.Fatal(err)
		}

		videoData, _ := c.FormFile("data")
		if videoData == nil {
			fmt.Println("the video data is nil")
		}
		fmt.Println("video data: ", videoData)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	// 登录接口
	NoAuthAPI.POST("/login", api.Login)
	// 注册接口
	NoAuthAPI.POST("/register", api.Register)
	// 获取验证码接口
	NoAuthAPI.POST("/register/vlcode", api.GetVerificationCodeHandler)

	err := r.Run(":8300")
	if err != nil {
		log.Fatal("http server run failed!")
	}
}
