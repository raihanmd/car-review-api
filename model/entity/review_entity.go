package entity

import "time"

type Review struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	CarID     uint   `gorm:"not null;index:idx_car_id_user_id,unique"`
	UserID    uint   `gorm:"not null;index:idx_car_id_user_id,unique"`
	Title     string `gorm:"not null;type:varchar(100)"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
	Car       Car  `gorm:"foreignKey:CarID"`
}
