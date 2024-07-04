package entity

import (
	"time"
)

var RoleAdmin = "ADMIN"
var RoleUser = "USER"

type Profile struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UserID    uint    `gorm:"unique;not null"`
	Email     *string `gorm:"type:varchar(100)"`
	FullName  *string `gorm:"type:varchar(100)"`
	Bio       *string `gorm:"type:text"`
	Age       *int    `gorm:"type:smallint"`
	Gender    *string `gorm:"type:varchar(6)" sql:"type:enum('MALE', 'FEMALE')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
