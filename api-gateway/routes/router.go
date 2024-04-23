package routes

import (
	"api-gateway/handler"
	"api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoute(m gin.HandlerFunc) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(m)
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

		// 注册
		v1.POST("/user/register", handler.UserRegister)
		v1.POST("/user/login", handler.UserLogin)

		task := v1.Group("/task")
		task.Use(middleware.JWT())
		task.POST("/create", handler.CreateTask)
		task.PUT("/update", handler.UpdateTask)
		task.DELETE("/delete", handler.DeleteTask)
		task.GET("/list", handler.TaskList)
	}
	return ginRouter
}
