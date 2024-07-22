package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"unique;not null;type:varchar(20)"`
	Email     string `gorm:"unique;not null;type:varchar(50)"`
	Password  string `gorm:"not null"`
	Role      string `sql:"type:enum('ADMIN', 'USER')" gorm:"default:'USER'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   Profile `gorm:"foreignKey:UserID"`
}
