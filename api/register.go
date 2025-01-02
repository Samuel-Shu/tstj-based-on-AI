package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tstj-based-on-AI/model"
)

func Register(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		log.Fatal("[api-Register]获取用户名或密码失败！", err)
	}

	register := model.Register(user.Name, user.Password, user.Email)

	if register == 1 {
		ctx.JSON(http.StatusOK,
			model.HttpStatus{
				StatusCode: 200,
				StatusMsg:  "用户名已存在，请做出修改后重试",
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
