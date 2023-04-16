package handler

import (
	"todolist/service"

	"github.com/labstack/echo"
)

type TaskHandler interface {
	GetTasks(ctx echo.Context) error
	CreateTask(ctx echo.Context) error
	UpdateTask(ctx echo.Context) error
	DeleteTask(ctx echo.Context) error
}

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return taskHandler{taskService}
}

func (h taskHandler) GetTasks(ctx echo.Context) error {
	return nil
}

func (h taskHandler) CreateTask(ctx echo.Context) error {
	return nil
}

func (h taskHandler) UpdateTask(ctx echo.Context) error {
	return nil
}

func (h taskHandler) DeleteTask(ctx echo.Context) error {
	return nil
}
