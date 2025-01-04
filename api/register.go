package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"tstj-based-on-AI/db"
	"tstj-based-on-AI/model"
	"tstj-based-on-AI/utils"
)

func Register(ctx *gin.Context) {
	type R struct {
		model.User
		VerificationCode string `json:"verificationCode"`
	}
	r := R{}
	if err := ctx.ShouldBind(&r); err != nil {
		log.Fatal("[api-Register]获取用户名或密码失败！", err)
	}
	fmt.Println("log:", r.Name, r.Password, r.Email)

	if r.Name == "" || r.Password == "" || r.Email == "" || r.VerificationCode == "" {
		ctx.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 400,
			StatusMsg:  "请输入必填项",
		})
		return
	}

	code := db.Db.Redis.Get(context.Background(), r.Email).Val()
	fmt.Println(code)
	if code != r.VerificationCode {
		ctx.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 400,
			StatusMsg:  "验证码错误或失效，请重试",
		})
		return
	}
	register := model.Register(r.Name, r.Password, r.Email)

	if register == 1 {
		ctx.JSON(http.StatusOK,
			model.HttpStatus{
				StatusCode: 400,
				StatusMsg:  "用户名已存在，请做出修改后重试",
			})
		return
	}

	ctx.JSON(http.StatusOK, model.HttpStatus{
		StatusCode: 200,
		StatusMsg:  "注册成功，请登录",
	})
}

func GetVerificationCodeHandler(ctx *gin.Context) {
	type E struct {
		Email string `json:"email"`
	}
	e := E{}
	if err := ctx.ShouldBind(&e); err != nil {
		ctx.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 400,
			StatusMsg:  "请输入正确邮箱",
		})
		log.Fatal("[api-GetVerificationCodeHandler]获取邮箱失败")
	}

	if e.Email == "" {
		ctx.JSON(http.StatusOK, model.HttpStatus{
			StatusCode: 400,
			StatusMsg:  "请输入正确邮箱",
		})
		return
	}
	fmt.Println(e.Email)
	vlcode := utils.CreateCaptcha()
	go func() {
		utils.Email(e.Email, vlcode)
	}()
	db.Db.Redis.Set(context.Background(), e.Email, vlcode, 60*time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"status_msg":  "发送验证码成功",
	})
}
