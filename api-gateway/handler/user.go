package handler

import (
	"api-gateway/service"
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Token    string `json:"token"`
	UserID   uint32 `json:"userId"`
	UserName string `json:"username"`
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

	token, err := utils.GenerateToken(uint(u.Data.UserID))

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "messge": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"user": UserResponse{
			Token:    token,
			UserID:   u.Data.UserID,
			UserName: u.Data.UserName,
		},
	})
}
