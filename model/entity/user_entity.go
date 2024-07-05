package entity

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null;type:varchar(20)"`
	Password string `gorm:"not null"`
	Role     string `sql:"type:enum('ADMIN', 'USER')" gorm:"default:'USER'"`
	// Reviews   []Review  `gorm:"foreignKey:UserID;references:ID"`
	// Comments  []Comment `gorm:"foreignKey:UserID;references:ID"`
	// Ratings   []Rating  `gorm:"foreignKey:UserID;references:ID"`
	// Favorites []Car     `gorm:"many2many:favorites;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   Profile `gorm:"foreignKey:UserID"`
	// Review    Review  `gorm:"foreignKey:UserID"`
}
