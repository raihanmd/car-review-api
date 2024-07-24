package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"github.com/raihanmd/fp-superbootcamp-go/model/web"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/request"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(*gin.Context, *request.ReviewCreateRequest, uint) (*response.ReviewResponse, error)
	Update(*gin.Context, *request.ReviewUpdateRequest, uint, uint) (*response.FindReviewResponse, error)
	Delete(*gin.Context, uint, uint) error
	FindAll(*gin.Context, *request.ReviewQueryRequest, *web.PaginationRequest) (*[]response.FindReviewResponse, *web.Metadata, error)
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
		UserID:   userID,
		CarID:    reviewCreateReq.CarID,
		Title:    reviewCreateReq.Title,
		Content:  reviewCreateReq.Content,
		ImageUrl: reviewCreateReq.ImageUrl,
	}

	if err := db.Model(&entity.Review{}).Create(&newReview).Take(&responseReview, "car_id = ?", newReview.CarID).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// violation foreign key car_id
			if pgErr.Code == "23503" {
				return nil, exceptions.NewCustomError(http.StatusNotFound, "Car not found")
			}
			// Handle duplicate key error
			if pgErr.Code == "23505" {
				return nil, exceptions.NewCustomError(http.StatusConflict, "You have reviewed this car")
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

func (service *reviewServiceImpl) FindAll(c *gin.Context, reviewQueryReq *request.ReviewQueryRequest, paging *web.PaginationRequest) (*[]response.FindReviewResponse, *web.Metadata, error) {
	db, _ := helper.GetDBAndLogger(c)

	var reviews []map[string]interface{}

	query := db.Table("reviews").
		Order("created_at desc").
		Select("reviews.*, reviews.id as review_id, cars.id as car_id, users.username, users.id as user_id").
		Joins("left join cars on reviews.car_id = cars.id").
		Joins("left join users on reviews.user_id = users.id")

	if reviewQueryReq.Title != nil {
		query = query.Where("to_tsvector('english', reviews.title) @@ plainto_tsquery('english', ?)", *reviewQueryReq.Title)
	}

	if reviewQueryReq.CarID != nil {
		query = query.Where("reviews.car_id = ?", *reviewQueryReq.CarID)
	}

	query.Count(&paging.TotalData)

	offset := (paging.Page - 1) * paging.Limit
	query = query.Limit(paging.Limit).Offset(offset)

	if err := query.Find(&reviews).Error; err != nil {
		return nil, nil, err
	}

	paging.TotalPages = int((paging.TotalData + int64(paging.Limit) - 1) / int64(paging.Limit))

	var responseReviews []response.FindReviewResponse

	for _, v := range reviews {
		review := response.FindReviewResponse{
			ID:        uint(v["review_id"].(int64)),
			Title:     v["title"].(string),
			Content:   v["content"].(string),
			ImageUrl:  v["image_url"].(string),
			CreatedAt: v["created_at"].(time.Time),
			UpdatedAt: v["updated_at"].(time.Time),
			User: response.ReviewUserResponse{
				ID:       uint(v["user_id"].(int64)),
				Username: v["username"].(string),
			},
		}

		responseReviews = append(responseReviews, review)
	}

	metadata := web.Metadata{
		Page:       &paging.Page,
		Limit:      &paging.Limit,
		TotalPages: &paging.TotalPages,
		TotalData:  &paging.TotalData,
	}

	return &responseReviews, &metadata, nil
}

func (service *reviewServiceImpl) FindByID(c *gin.Context, reviewId uint) (*response.FindReviewResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var review map[string]any

	if err := db.Table("reviews").Select("reviews.*, reviews.id as review_id, cars.id as car_id, users.username, users.id as user_id").
		Joins("left join cars on reviews.car_id = cars.id").
		Joins("left join users on reviews.user_id = users.id").
		Take(&review, "reviews.id = ?", reviewId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewCustomError(http.StatusNotFound, "review not found")
		}
		return nil, err
	}

	responseReview := response.FindReviewResponse{
		ID:        uint(review["review_id"].(int64)),
		Title:     review["title"].(string),
		Content:   review["content"].(string),
		ImageUrl:  review["image_url"].(string),
		CreatedAt: review["created_at"].(time.Time),
		UpdatedAt: review["updated_at"].(time.Time),
		User: response.ReviewUserResponse{
			ID:       uint(review["user_id"].(int64)),
			Username: review["username"].(string),
		},
	}

	return &responseReview, nil
}
