package model

import (
	"errors"
	"fmt"
	"task/global"
	"task/service"
	"time"
)

type Task struct {
	ID        uint32     `gorm:"primarykey;unique;column:id"`
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

func (t *Task) UpdateTask(req *service.TaskRequest) (*Task, error) {
	if isExist := t.CheckTaskIsExist(req.TaskID); !isExist {
		return nil, errors.New("任务不存在")
	}
	task := Task{
		ID:        req.TaskID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	if err := global.DB.Model(&Task{}).Where("id = ?", req.TaskID).Updates(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (*Task) CheckTaskIsExist(taskID uint32) bool {
	var task Task
	if err := global.DB.Model(&Task{}).Where("id = ?", taskID).First(&task).Error; err != nil {
		return false
	}
	return true
}
