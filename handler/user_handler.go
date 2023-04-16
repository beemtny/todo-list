package handler

import (
	"todolist/service"

	"github.com/labstack/echo"
)

type UserHandler interface {
	Login(ctx echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandler{userService}
}

func (h userHandler) Login(ctx echo.Context) error {
	return nil
}
