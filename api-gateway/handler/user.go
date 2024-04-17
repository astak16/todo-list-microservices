package handler

import (
	"api-gateway/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {

	userName := ctx.Query("userName")
	password := ctx.Query("password")
	if userName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " username 不能为空"})
		return
	}
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "password 不能为空"})
		return
	}
	userReq := service.UserRequest{
		UserName: userName,
		Password: password,
	}

	fmt.Println("userReq", ctx.Keys["user"])

	userService := ctx.Keys["user"].(service.UserServiceClient)
	userService.Register(ctx, &userReq)
	ctx.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
