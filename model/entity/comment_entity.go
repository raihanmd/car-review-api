package entity

import "time"

type Comment struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	ReviewID uint `gorm:"not null"`
	// Review    Review
	UserID uint `gorm:"not null"`
	// User      User
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
