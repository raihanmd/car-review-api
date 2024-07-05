package entity

import "time"

type Car struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Brand string `gorm:"not null;type:varchar(50)"`
	Model string `gorm:"not null;type:varchar(50)"`
	Year  int    `gorm:"not null;type:smallint"`
	Image string `gorm:"not null"`
	// Reviews     []Review `gorm:"foreignKey:CarID"`
	// Ratings     []Rating `gorm:"foreignKey:CarID"`
	// FavoritedBy []User   `gorm:"many2many:favorites;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Review    Review `gorm:"foreignKey:CarID"`
}
