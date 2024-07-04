package entity

import "time"

type Favorite struct {
	UserID    uint `gorm:"primaryKey"`
	CarID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
