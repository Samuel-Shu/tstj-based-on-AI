package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
	"tstj-based-on-AI/config"
	"tstj-based-on-AI/model"
)

type MyClaim struct {
	UserId   int64
	UserName string
	jwt.RegisteredClaims
}

// GenerateToken Generate generate jwtToken
func GenerateToken(uerName string, userId int64, ctx *gin.Context) string {
	claim := MyClaim{
		UserId:   userId,
		UserName: uerName,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                    //effective time
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //sign time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)), //expire time
		},
	}
	//use HAS256 to sign a jwtToken with claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtKey, exist := ctx.Get("ServerConfig")
	if !exist {
		log.Fatal(errors.New("①get JwtKey failed"))
	}
	//fmt.Println(jwtKey.(*config.ServerConfig).JwtKey)
	jwtToken, err := token.SignedString([]byte(jwtKey.(*config.ServerConfig).JwtKey))
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(jwtToken)
	return jwtToken
}

// ParseToken  can parse jwtToken to get Claim's information
func ParseToken(jwtToken string, ctx *gin.Context) (any, error) {
	claims := MyClaim{}

	j, exist := ctx.Get("ServerConfig")
	if !exist {
		log.Fatal(errors.New("②get JwtKey failed"))
	}
	//fmt.Println(jwtToken)
	if jwtToken == "" {
		return nil, errors.New("you don't have valid token")
	}
	_, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("your alg is not correct,your alg is %v", token.Header["alg"])
		}
		return []byte(j.(*config.ServerConfig).JwtKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return claims, err
	}
	return claims, nil
}

// JWT gin中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		type TokenData struct {
			Token string `json:"token" form:"token" binding:"required"`
		}
		var tokenData TokenData
		err := c.ShouldBind(&tokenData)
		if err != nil {
			fmt.Println(err)
		}

		//token为空情况，返回错误代码
		if tokenData.Token == "" {
			c.JSON(http.StatusOK, model.HttpStatus{
				StatusCode: -1,
				StatusMsg:  "token is not exist",
			},
			)
			c.Abort()
			return
		}

		parseToken, err1 := ParseToken(tokenData.Token, c)
		//token解析错误情况，返回错误代码
		if err1 != nil {
			fmt.Println(err1)
			c.JSON(http.StatusOK, model.HttpStatus{
				StatusCode: -1,
				StatusMsg:  "1：token is incorrect or invalid",
			})
			c.Abort()
			return
		}
		//验证token对应的用户以及被移但token仍存在的情况
		if !model.FindUserWithId(parseToken.(MyClaim).UserId) {
			c.JSON(http.StatusOK, model.HttpStatus{
				StatusCode: -1,
				StatusMsg:  "user is not exit or token is invalid",
			})
			c.Abort()
			return
		}
		//判断token与用户id是否匹配
		//todo

		//成功通过验证，将userId写入gin.context上下文，便于后续使用
		c.Set("userId", parseToken.(MyClaim).UserId)
		c.Next()
	}
}
