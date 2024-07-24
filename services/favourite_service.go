package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"go.uber.org/zap"
)

type FavouriteService interface {
	FavouriteCar(c *gin.Context, carID uint, userID uint) error
	UnfavouriteCar(c *gin.Context, carID uint, userID uint) error
	GetUserFavourites(c *gin.Context, userID uint) (*[]response.FavouriteResponse, error)
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
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// unique violation
			if pgErr.Code == "23505" {
				return exceptions.NewCustomError(http.StatusBadRequest, "You have favourited this car")
			}
			// violation foreign key car_id
			if pgErr.Code == "23503" {
				return exceptions.NewCustomError(http.StatusNotFound, "Car not found")
			}
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

func (service *favouriteServiceImpl) GetUserFavourites(c *gin.Context, userID uint) (*[]response.FavouriteResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var favouriteResponses []response.FavouriteResponse

	if err := db.Model(&entity.Favourite{}).
		Select("cars.id as car_id, brands.name as brand, cars.model, cars.image_url").
		Joins("left join cars on favourites.car_id = cars.id").
		Joins("left join brands on cars.brand_id = brands.id").
		Where("user_id = ?", userID).
		Find(&favouriteResponses).Error; err != nil {
		return nil, err
	}

	return &favouriteResponses, nil
}
