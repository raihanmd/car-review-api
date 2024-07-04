package entity

import "time"

type Review struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	CarID uint `gorm:"not null"`
	// Car     Car
	UserID uint `gorm:"not null"`
	// User    User
	Title   string `gorm:"not null;type:varchar(100)"`
	Content string `gorm:"not null"`
	// Comments  []Comment `gorm:"foreignKey:ReviewID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
