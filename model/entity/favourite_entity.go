package entity

type Favourite struct {
	UserID uint `gorm:"index:idx_car_id_user_id,unique"`
	CarID  uint `gorm:"index:idx_car_id_user_id,unique"`
}
