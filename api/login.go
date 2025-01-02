package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tstj-based-on-AI/model"
	"tstj-based-on-AI/utils"
)

func Login(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		log.Fatal("[api-Login]获取用户名或密码失败！", err)
	}

	if !model.Login(user.Name, utils.Md5(user.Password)) {
		ctx.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 400,
			StatusMsg:  "用户名或密码错误",
		})
		return
	}

	userInfo := model.FindUserWithName(user.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"name": userInfo.Name,
		"role": userInfo.Role,
	})
}
