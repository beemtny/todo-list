package entity

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"username"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"autoUpdateTime;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime;default:current_timestamp"`
}
