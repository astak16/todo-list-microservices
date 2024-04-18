package handler

import (
	"api-gateway/service"
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

	userService := ctx.Keys["user"].(service.UserServiceClient)
	u, err := userService.Register(ctx, &userReq)
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
