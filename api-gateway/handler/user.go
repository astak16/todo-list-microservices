package handler

import (
	"api-gateway/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRegister(ctx *gin.Context) {
	userReq := User{}
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}

	userService := ctx.Keys["user"].(service.UserServiceClient)
	u, err := userService.Register(ctx, &service.UserRequest{
		UserName: userReq.UserName,
		Password: userReq.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
		"user": service.UserModel{
			UserID:   u.Data.UserID,
			UserName: u.Data.UserName,
		},
	})
}

func UserLogin(ctx *gin.Context) {
	userReq := User{}
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}

	userService := ctx.Keys["user"].(service.UserServiceClient)
	u, err := userService.Login(ctx, &service.UserRequest{
		UserName: userReq.UserName,
		Password: userReq.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}
	fmt.Println(u, "uuuussss")

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"user": service.UserModel{
			UserID:   u.Data.UserID,
			UserName: u.Data.UserName,
		},
	})
}
