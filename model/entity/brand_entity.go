package entity

type Brand struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;type:varchar(50)"`
}
