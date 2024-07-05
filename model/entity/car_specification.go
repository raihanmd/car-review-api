package entity

type CarSpecification struct {
	ID             uint              `gorm:"primaryKey;autoIncrement"`
	CarID          uint              `gorm:"not null"`
	Dimension      CarDimension      `gorm:"embedded"`
	Engine         string            `gorm:"not null;type:varchar(200)"`
	Torque         int16             `gorm:"not null;type:smallint"`
	Transmission   string            `gorm:"not null;type:varchar(50)"`
	Acceleration   float32           `gorm:"not null;type:real"`
	HorsePower     int16             `gorm:"not null;type:smallint"`
	BreakingSystem CarBreakingSystem `gorm:"embedded;embeddedPrefix:breaking_system_"`
	Fuel           string            `gorm:"not null;type:varchar(50)"`
}

type CarDimension struct {
	Width  int16 `gorm:"not null;type:smallint"`
	Height int16 `gorm:"not null;type:smallint"`
	Length int16 `gorm:"not null;type:smallint"`
}

type CarBreakingSystem struct {
	Front string `gorm:"not null;type:varchar(50)"`
	Back  string `gorm:"not null;type:varchar(50)"`
}
