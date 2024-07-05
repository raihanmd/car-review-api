package entity

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	ReviewID  uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Review    Review `gorm:"foreignKey:ReviewID"`
	User      User   `gorm:"foreignKey:UserID"`
}
