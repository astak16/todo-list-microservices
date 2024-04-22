package handler

import (
	"api-gateway/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskCreate struct {
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type TaskUpdate struct {
	ID        uint32 `json:"id" binding:"required"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

func CreateTask(ctx *gin.Context) {
	var task TaskCreate
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

func UpdateTask(ctx *gin.Context) {
	var task TaskUpdate
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}
	taskService := ctx.Keys["task"].(service.TaskServiceClient)
	t, err := taskService.TaskUpdate(ctx, &service.TaskRequest{
		TaskID:    task.ID,
		Status:    task.Status,
		Title:     task.Title,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 400, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
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
