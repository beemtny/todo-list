package entity

import "time"

type Tasks struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"user_id"`
	Description string    `gorm:"description"`
	IsComplete  bool      `gorm:"is_complete"`
	CreatedAt   time.Time `gorm:"autoUpdateTime;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;default:current_timestamp"`
	Users       Users     `gorm:"foreignKey:UserID"`
}
