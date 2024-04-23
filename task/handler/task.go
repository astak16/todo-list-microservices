package handler

import (
	"context"
	"task/model"
	"task/service"
)

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (*TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.TaskResponse, err error) {
	var task model.Task
	t, err := task.CreateTask(req)
	if err != nil {
		return nil, err
	}

	return &service.TaskResponse{
		Code: 200,
		TaskDetail: &service.TaskModel{
			TaskID:    uint32(t.ID),
			Status:    t.Status,
			Title:     t.Title,
			Content:   t.Content,
			StartTime: t.StartTime,
			EndTime:   t.EndTime,
		},
	}, nil
}

func (*TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.TaskResponse, err error) {
	var task model.Task
	t, err := task.UpdateTask(req)
	if err != nil {
		return nil, err
	}
	return &service.TaskResponse{
		Code: 200,
		TaskDetail: &service.TaskModel{
			TaskID:    t.ID,
			Status:    t.Status,
			Title:     t.Title,
			Content:   t.Content,
			StartTime: t.StartTime,
			EndTime:   t.EndTime,
		},
	}, nil
}

func (*TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.TaskResponse, err error) {
	var task model.Task
	_, err = task.DeleteTask(req)
	if err != nil {
		return nil, err
	}
	return &service.TaskResponse{
		Code: 200,
	}, nil
}

func (t *TaskService) TaskDetail(ctx context.Context, req *service.TaskRequest) (resp *service.TaskResponse, err error) {
	return nil, nil
}

func (t *TaskService) TaskList(ctx context.Context, req *service.TaskRequest) (resp *service.TaskListResponse, err error) {
	var task model.Task
	tasks, err := task.TaskList(req)
	if err != nil {
		return nil, err
	}
	var taskList []*service.TaskModel
	for _, t := range tasks {
		taskList = append(taskList, &service.TaskModel{
			TaskID:    t.ID,
			Status:    t.Status,
			Title:     t.Title,
			Content:   t.Content,
			StartTime: t.StartTime,
			EndTime:   t.EndTime,
		})
	}
	return &service.TaskListResponse{
		Code:  200,
		Tasks: taskList,
	}, nil
}
