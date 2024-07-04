package app

import (
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(helper.MustGetEnv("DB_DSN")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Car{}, &entity.Review{}, &entity.Comment{}, &entity.Rating{}, &entity.Favorite{}, &entity.Profile{})
	helper.PanicIfError(err)

	return db
}
