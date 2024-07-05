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

	newCar := service.toCarEntity(carCreateReq)

	if err := db.Create(newCar).Error; err != nil {
		return nil, err
	}

	logger.Info("car created successfully", zap.Uint("carID", newCar.ID))

	return service.toCarResponse(newCar), nil
}

func (service *carServiceImpl) Update(c *gin.Context, carUpdateReq *request.CarUpdateRequest, carID uint) (*response.CarResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	updateCar := service.toCarEntity(carUpdateReq)

	var responseCar entity.Car

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.Model(&entity.Car{}).Where("id = ?", carID).Updates(updateCar)

		if result.RowsAffected == 0 {
			return exceptions.NewCustomError(http.StatusNotFound, "car not found")
		}

		if result.Error != nil {
			return result.Error
		}

		if err := db.Model(&entity.CarSpecification{CarID: carID}).Where("car_id = ?", carID).Updates(updateCar.CarSpecification).Error; err != nil {
			return err
		}

		if err := tx.Preload("CarSpecification").Take(&responseCar, "id = ?", carID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("car updated successfully", zap.Uint("carID", carID))

	return service.toCarResponse(&responseCar), nil
}

func (service *carServiceImpl) Delete(c *gin.Context, carID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.Model(&entity.CarSpecification{}).Where("car_id = ?", carID).Delete(&entity.CarSpecification{})

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return exceptions.NewCustomError(http.StatusNotFound, "car not found")
		}

		err := db.Where("id = ?", carID).Delete(&entity.Car{}).Error

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("car deleted successfully", zap.Uint("carID", carID))

	return nil
}

func (service *carServiceImpl) FindAll(c *gin.Context) (*[]response.CarResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var cars []entity.Car

	if err := db.Preload("CarSpecification").Find(&cars).Error; err != nil {
		return nil, err
	}

	var responseCars []response.CarResponse

	for _, car := range cars {
		responseCars = append(responseCars, *service.toCarResponse(&car))
	}

	return &responseCars, nil
}

func (service *carServiceImpl) FindByID(c *gin.Context, carId uint) (*response.CarResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var car entity.Car

	if err := db.Preload("CarSpecification").Take(&car, "id = ?", carId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "car not found")
		}
		return nil, err
	}

	return service.toCarResponse(&car), nil
}

func (service *carServiceImpl) toCarResponse(car *entity.Car) *response.CarResponse {
	return &response.CarResponse{
		ID:                  car.ID,
		BrandID:             car.BrandID,
		Model:               car.Model,
		Year:                car.Year,
		ImageUrl:            car.ImageUrl,
		Width:               car.CarSpecification.Dimension.Width,
		Height:              car.CarSpecification.Dimension.Height,
		Length:              car.CarSpecification.Dimension.Length,
		Engine:              car.CarSpecification.Engine,
		Torque:              car.CarSpecification.Torque,
		Transmission:        car.CarSpecification.Transmission,
		Acceleration:        car.CarSpecification.Acceleration,
		HorsePower:          car.CarSpecification.HorsePower,
		BreakingSystemFront: car.CarSpecification.BreakingSystem.Front,
		BreakingSystemBack:  car.CarSpecification.BreakingSystem.Back,
		Fuel:                car.CarSpecification.Fuel,
	}
}

func (service *carServiceImpl) toCarEntity(req any) *entity.Car {
	switch r := req.(type) {
	case *request.CarCreateRequest:
		return &entity.Car{
			BrandID:  r.BrandID,
			Model:    r.Model,
			Year:     r.Year,
			ImageUrl: r.ImageUrl,
			CarSpecification: entity.CarSpecification{
				Dimension: entity.CarDimension{
					Length: r.Length,
					Width:  r.Width,
					Height: r.Height,
				},
				Engine:       r.Engine,
				Torque:       r.Torque,
				Transmission: r.Transmission,
				Acceleration: r.Acceleration,
				HorsePower:   r.HorsePower,
				BreakingSystem: entity.CarBreakingSystem{
					Front: r.BreakingSystemFront,
					Back:  r.BreakingSystemBack,
				},
				Fuel: r.Fuel,
			},
		}
	case *request.CarUpdateRequest:
		return &entity.Car{
			BrandID:  r.BrandID,
			Model:    r.Model,
			Year:     r.Year,
			ImageUrl: r.ImageUrl,
			CarSpecification: entity.CarSpecification{
				Dimension: entity.CarDimension{
					Length: r.Length,
					Width:  r.Width,
					Height: r.Height,
				},
				Engine:       r.Engine,
				Torque:       r.Torque,
				Transmission: r.Transmission,
				Acceleration: r.Acceleration,
				HorsePower:   r.HorsePower,
				BreakingSystem: entity.CarBreakingSystem{
					Front: r.BreakingSystemFront,
					Back:  r.BreakingSystemBack,
				},
				Fuel: r.Fuel,
			},
		}
	default:
		return nil
	}
}
