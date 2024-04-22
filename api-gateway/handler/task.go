package handler

import (
	"api-gateway/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

func CreateTask(ctx *gin.Context) {

	var task Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}

	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	t, err := taskService.TaskCreate(ctx, &service.TaskRequest{
		Status:    task.Status,
		Title:     task.Title,
		Content:   task.Content,
		UserID:    "22",
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"task": service.TaskModel{
			TaskID:    t.TaskDetail.TaskID,
			Status:    t.TaskDetail.Status,
			Title:     t.TaskDetail.Title,
			Content:   t.TaskDetail.Content,
			StartTime: t.TaskDetail.StartTime,
			EndTime:   t.TaskDetail.EndTime,
		},
	})
}
