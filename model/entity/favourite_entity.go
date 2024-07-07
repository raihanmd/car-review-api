package entity

type Favourite struct {
	UserID uint `gorm:"index:idx_favourite,unique"`
	CarID  uint `gorm:"index:idx_favourite,unique"`
	User   User `gorm:"foreignKey:UserID"`
	Car    Car  `gorm:"foreignKey:CarID"`
}
