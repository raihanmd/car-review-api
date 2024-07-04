package entity

import "time"

type Rating struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	CarID uint `gorm:"not null"`
	// Car       Car
	UserID uint `gorm:"not null"`
	// User      User
	Rating    int `gorm:"not null;type:smallint"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
