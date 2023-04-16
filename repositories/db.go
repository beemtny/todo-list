package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := "host=localhost user=root password=P@ssw0rd dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
