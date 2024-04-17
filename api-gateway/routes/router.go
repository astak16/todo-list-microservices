package routes

import (
	"api-gateway/handler"

	"github.com/gin-gonic/gin"
)

func NewRoute(middleware gin.HandlerFunc) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware)
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

		// 注册
		v1.POST("/user/register", handler.UserRegister)
	}
	return ginRouter
}
