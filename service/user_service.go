package service

import "todolist/repositories"

type UserService interface {
}

type userService struct {
	userRepo repositories.UsersRepository
}

func NewUserService(userRepo repositories.UsersRepository) UserService {
	return userService{userRepo}
}
