package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
	"github.com/raihanmd/car-review-sb/model/web/request"
	"github.com/raihanmd/car-review-sb/model/web/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CarService interface {
	Create(*gin.Context, *request.CarCreateRequest) (*response.CarResponse, error)
	Update(*gin.Context, *request.CarUpdateRequest, uint) (*response.CarResponse, error)
	Delete(*gin.Context, uint) error
	FindAll(*gin.Context) (*[]response.CarResponse, error)
	FindByID(*gin.Context, uint) (*response.CarResponse, error)
}

type carServiceImpl struct{}

func NewCarService() CarService {
	return &carServiceImpl{}
}

func (service *carServiceImpl) Create(c *gin.Context, carCreateReq *request.CarCreateRequest) (*response.CarResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var responseCar response.CarResponse

	newCar := entity.Car{
		Brand: carCreateReq.Brand,
		Model: carCreateReq.Model,
		Year:  carCreateReq.Year,
		Image: carCreateReq.Image,
	}

	if err := db.Model(&entity.Car{}).Create(&newCar).Take(&responseCar, "id = ?", newCar.ID).Error; err != nil {
		return nil, err
	}

	logger.Info("car created successfully", zap.Uint("carID", responseCar.ID))

	return &responseCar, nil
}

func (service *carServiceImpl) Update(c *gin.Context, carUpdateReq *request.CarUpdateRequest, carID uint) (*response.CarResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var responseCar response.CarResponse

	result := db.Model(&entity.Car{}).Where("id = ?", carID).Updates(carUpdateReq).Take(&responseCar, "id = ?", carID)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, exceptions.NewCustomError(http.StatusNotFound, "car not found")
	}

	logger.Info("car updated successfully", zap.Uint("carID", carID))

	return &responseCar, nil
}

func (service *carServiceImpl) Delete(c *gin.Context, carID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	result := db.Where("id = ?", carID).Delete(&entity.Car{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exceptions.NewCustomError(http.StatusNotFound, "car not found")
	}

	logger.Info("car deleted successfully", zap.Uint("carID", carID))

	return nil
}

func (service *carServiceImpl) FindAll(c *gin.Context) (*[]response.CarResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var responseCar []response.CarResponse

	if err := db.Model(&entity.Car{}).Find(&responseCar).Error; err != nil {
		return nil, err
	}

	return &responseCar, nil
}

func (service *carServiceImpl) FindByID(c *gin.Context, carId uint) (*response.CarResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var responseCar response.CarResponse

	if err := db.Model(&entity.Car{}).Take(&responseCar, "id = ?", carId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "car not found")
		}
		return nil, err
	}

	return &responseCar, nil
}
