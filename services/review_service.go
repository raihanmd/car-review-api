package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
	"github.com/raihanmd/car-review-sb/model/web/request"
	"github.com/raihanmd/car-review-sb/model/web/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(*gin.Context, *request.ReviewCreateRequest, uint) (*response.ReviewResponse, error)
	Update(*gin.Context, *request.ReviewUpdateRequest, uint, uint) (*response.FindReviewResponse, error)
	Delete(*gin.Context, uint, uint) error
	FindAll(*gin.Context) (*[]response.FindReviewResponse, error)
	FindByID(*gin.Context, uint) (*response.FindReviewResponse, error)
}

type reviewServiceImpl struct{}

func NewreviewService() ReviewService {
	return &reviewServiceImpl{}
}

func (service *reviewServiceImpl) Create(c *gin.Context, reviewCreateReq *request.ReviewCreateRequest, userID uint) (*response.ReviewResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var responseReview response.ReviewResponse

	newReview := entity.Review{
		UserID:  userID,
		CarID:   reviewCreateReq.CarID,
		Title:   reviewCreateReq.Title,
		Content: reviewCreateReq.Content,
	}

	if err := db.Model(&entity.Review{}).Create(&newReview).Take(&responseReview, "car_id = ?", newReview.CarID).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// violation foreign key car_id
			if pgErr.Code == "23503" {
				return nil, exceptions.NewCustomError(http.StatusNotFound, "car not found")
			}
			// Handle duplicate key error
			if pgErr.Code == "23505" {
				return nil, exceptions.NewCustomError(http.StatusConflict, "only one review for this car permitted")
			}
		}
		return nil, err
	}

	logger.Info("review created successfully", zap.Uint("reviewID", responseReview.ID), zap.Any("userID", userID))

	return &responseReview, nil
}

func (service *reviewServiceImpl) Update(c *gin.Context, reviewUpdateReq *request.ReviewUpdateRequest, userID, reviewID uint) (*response.FindReviewResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	responseReview, err := service.FindByID(c, reviewID)

	if err != nil {
		return nil, err
	}

	result := db.Model(&entity.Review{}).
		Where("id = ?", reviewID).
		Where("user_id = ?", userID).
		Where("car_id = ?", responseReview.Car.ID).
		Updates(reviewUpdateReq).
		Take(&responseReview, "id = ?", reviewID)

	if result.Error != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, exceptions.NewCustomError(http.StatusNotFound, "review not found")
	}

	logger.Info("review created successfully", zap.Uint("reviewID", responseReview.ID), zap.Any("userID", userID))

	return responseReview, nil
}

func (service *reviewServiceImpl) Delete(c *gin.Context, userID, reviewID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	result := db.Where("id = ?", reviewID).Where("user_id = ?", userID).Delete(&entity.Review{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exceptions.NewCustomError(http.StatusNotFound, "review not found")
	}

	logger.Info("review deleted successfully", zap.Uint("reviewID", reviewID))

	return nil
}

func (service *reviewServiceImpl) FindAll(c *gin.Context) (*[]response.FindReviewResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var reviews []map[string]any

	if err := db.Table("reviews").Select("reviews.*, reviews.id as review_id, cars.*, cars.id as car_id, users.username, users.id as user_id").
		Joins("left join cars on reviews.car_id = cars.id").
		Joins("left join users on reviews.user_id = users.id").
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	var responseReviews []response.FindReviewResponse

	for _, v := range reviews {
		review := response.FindReviewResponse{
			ID:        v["review_id"].(uint),
			Title:     v["title"].(string),
			Content:   v["content"].(string),
			Image:     v["image"].(string),
			CreatedAt: v["created_at"].(time.Time),
			UpdatedAt: v["updated_at"].(time.Time),
			Car: response.CarResponse{
				ID:       v["car_id"].(uint),
				BrandID:  v["brand_id"].(uint),
				Model:    v["model"].(string),
				Year:     v["year"].(int16),
				ImageUrl: v["image_url"].(string),
			},
			User: response.ReviewUserResponse{
				ID:       uint(v["user_id"].(int64)),
				Username: v["username"].(string),
			},
		}

		responseReviews = append(responseReviews, review)
	}

	return &responseReviews, nil
}

func (service *reviewServiceImpl) FindByID(c *gin.Context, reviewId uint) (*response.FindReviewResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var review map[string]any

	if err := db.Table("reviews").Select("reviews.*, reviews.id as review_id, cars.*, cars.id as car_id, users.username, users.id as user_id").
		Joins("left join cars on reviews.car_id = cars.id").
		Joins("left join users on reviews.user_id = users.id").
		Take(&review, "reviews.id = ?", reviewId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "reviewnot found")
		}
		return nil, err
	}

	return &response.FindReviewResponse{
		ID:        review["review_id"].(uint),
		Title:     review["title"].(string),
		Content:   review["content"].(string),
		Image:     review["image"].(string),
		CreatedAt: review["created_at"].(time.Time),
		UpdatedAt: review["updated_at"].(time.Time),
		Car: response.CarResponse{
			ID:       review["car_id"].(uint),
			BrandID:  review["brand_id"].(uint),
			Model:    review["model"].(string),
			Year:     review["year"].(int16),
			ImageUrl: review["image_url"].(string),
		},
		User: response.ReviewUserResponse{
			ID:       review["user_id"].(uint),
			Username: review["username"].(string),
		},
	}, nil
}
