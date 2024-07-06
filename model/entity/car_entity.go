package entity

import "time"

type Car struct {
	ID               uint             `gorm:"primaryKey;autoIncrement"`
	BrandID          uint             `gorm:"not null"`
	Model            string           `gorm:"not null;type:varchar(50)"`
	Year             int16            `gorm:"not null;type:smallint"`
	ImageUrl         string           `gorm:"not null;type:varchar(255)"`
	CarSpecification CarSpecification `gorm:"foreignKey:CarID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Brand            Brand `gorm:"foreignKey:BrandID"`
}
