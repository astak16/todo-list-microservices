package model

import (
	"fmt"
	"task/global"
	"task/service"
	"time"
)

type Task struct {
	ID        uint       `gorm:"primarykey;unique;column:id"`
	UserID    string     `gorm:"column:user_id;not null"`
	Status    string     `gorm:"column:status"`
	Title     string     `gorm:"column:title"`
	Content   string     `gorm:"column:content"`
	StartTime string     `gorm:"column:start_time"`
	EndTime   string     `gorm:"column:end_time"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}

func InitTask() {
	migartion()
}

func migartion() {
	global.DB.Set("gorm:table_options", fmt.Sprintf("charset=%s", global.MySqlConfig.Charset)).AutoMigrate(&Task{})
}

func (*Task) CreateTask(req *service.TaskRequest) (*Task, error) {
	task := Task{
		UserID:    req.UserID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	if err := global.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
