package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/request"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BrandService interface {
	Create(*gin.Context, *request.BrandRequest) (*response.BrandResponse, error)
	Update(*gin.Context, *request.BrandRequest, uint) (*response.BrandResponse, error)
	Delete(*gin.Context, uint) error
	FindAll(*gin.Context) (*[]response.BrandResponse, error)
}

type brandServiceImpl struct{}

func NewBrandService() BrandService {
	return &brandServiceImpl{}
}

func (service *brandServiceImpl) Create(c *gin.Context, brandCreateRequest *request.BrandRequest) (*response.BrandResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	newBrand := service.toBrandEntity(brandCreateRequest)

	if err := db.Create(newBrand).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// Handle duplicate key error
			if pgErr.Code == "23505" {
				return nil, exceptions.NewCustomError(http.StatusConflict, "Name already exist")
			}
		}
		return nil, err
	}

	logger.Info("brand created successfully", zap.Uint("brandID", newBrand.ID))

	return service.toBrandResponse(newBrand), nil
}

func (service *brandServiceImpl) Update(c *gin.Context, brandRequest *request.BrandRequest, brandID uint) (*response.BrandResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	updateBrand := service.toBrandEntity(brandRequest)

	var brand entity.Brand

	err := db.Transaction(func(tx *gorm.DB) error {
		result := db.Model(&entity.Brand{}).Where("id = ?", brandID).Updates(updateBrand)

		if result.Error != nil {
			if pgErr, ok := result.Error.(*pgconn.PgError); ok {
				// Handle duplicate key error
				if pgErr.Code == "23505" {
					return exceptions.NewCustomError(http.StatusConflict, "Name already exist")
				}
			}
			return result.Error
		}

		if result.RowsAffected == 0 {
			return exceptions.NewCustomError(http.StatusNotFound, "Brand not found")
		}

		if err := tx.Take(&brand, "id = ?", brandID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("brand updated successfully", zap.Uint("brandID", brandID))

	return service.toBrandResponse(&brand), nil
}

func (service *brandServiceImpl) Delete(c *gin.Context, brandID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	result := db.Where("id = ?", brandID).Delete(&entity.Brand{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exceptions.NewCustomError(http.StatusNotFound, "Brand not found")
	}

	logger.Info("brand deleted successfully", zap.Uint("brandID", brandID))

	return nil
}

func (service *brandServiceImpl) FindAll(c *gin.Context) (*[]response.BrandResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var brands []entity.Brand

	if err := db.Order("id").Find(&brands).Error; err != nil {
		return nil, err
	}

	var responseBrands []response.BrandResponse
	for _, brand := range brands {
		responseBrands = append(responseBrands, *service.toBrandResponse(&brand))
	}

	return &responseBrands, nil
}

func (service *brandServiceImpl) toBrandEntity(req *request.BrandRequest) *entity.Brand {
	return &entity.Brand{
		Name: req.Name,
	}
}

func (service *brandServiceImpl) toBrandResponse(brand *entity.Brand) *response.BrandResponse {
	return &response.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
	}
}
