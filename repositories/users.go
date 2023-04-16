package repositories

import (
	"todolist/repositories/entity"

	"gorm.io/gorm"
)

type UsersRepository interface {
	// Login()
}

type usersRepositoryDB struct {
	db *gorm.DB
}

func NewUsersRepositoryDB(db *gorm.DB) UsersRepository {
	db.AutoMigrate(&entity.Users{})

	return usersRepositoryDB{db}
}
