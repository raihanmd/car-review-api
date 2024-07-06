package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
	"go.uber.org/zap"
)

type FavouriteService interface {
	FavouriteCar(c *gin.Context, carID uint, userID uint) error
	UnfavouriteCar(c *gin.Context, carID uint, userID uint) error
}

type favouriteServiceImpl struct{}

func NewFavouriteService() FavouriteService {
	return &favouriteServiceImpl{}
}

func (service *favouriteServiceImpl) FavouriteCar(c *gin.Context, carID uint, userID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	favourite := entity.Favourite{
		UserID: userID,
		CarID:  carID,
	}

	if err := db.Create(&favourite).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // unique_violation
			return errors.New("car is already favourited by the user")
		}
		return err
	}

	logger.Info("car favourited successfully", zap.Uint("carID", carID), zap.Uint("userID", userID))

	return nil
}

func (service *favouriteServiceImpl) UnfavouriteCar(c *gin.Context, carID uint, userID uint) error {
	db, _ := helper.GetDBAndLogger(c)

	result := db.Where("user_id = ? AND car_id = ?", userID, carID).Delete(&entity.Favourite{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("favourite not found")
	}

	return nil
}
