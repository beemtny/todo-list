package service

import "todolist/repositories"

type TaskService interface {
}

type taskService struct {
	taskRepo repositories.TasksRepository
}

func NewTaskService(taskRepo repositories.TasksRepository) TaskService {
	return taskService{taskRepo}
}
