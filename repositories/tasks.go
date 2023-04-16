package repositories

import (
	"todolist/repositories/entity"

	"gorm.io/gorm"
)

type TasksRepository interface {
}

type tasksRepositoryDB struct {
	db *gorm.DB
}

func NewTasksRepositoryDB(db *gorm.DB) TasksRepository {
	db.AutoMigrate(&entity.Tasks{})

	return tasksRepositoryDB{db}
}

// func (t taskRepositoryDB)
