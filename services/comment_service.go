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

type CommentService interface {
	Create(*gin.Context, *request.CommentCreateRequest, uint) (*response.CommentResponse, error)
	Update(*gin.Context, *request.CommentUpdateRequest, uint, uint) (*response.CommentResponse, error)
	Delete(*gin.Context, uint, uint) error
	FindByReviewId(*gin.Context, uint) (*[]response.CommentResponse, error)
}

type commentServiceImpl struct{}

func NewCommentService() CommentService {
	return &commentServiceImpl{}
}

func (service *commentServiceImpl) Create(c *gin.Context, commentCreateReq *request.CommentCreateRequest, userID uint) (*response.CommentResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	newComment := service.toCommentEntity(commentCreateReq)

	newComment.UserID = userID

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newComment).Error; err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				// violation foreign key review_id
				if pgErr.Code == "23503" {
					return exceptions.NewCustomError(http.StatusNotFound, "Review not found")
				}
			}
			return err
		}

		if err := tx.Model(&entity.Comment{}).
			Preload("User", func(tx *gorm.DB) *gorm.DB {
				return tx.Select("id, username")
			}).
			Take(newComment).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("Comment created successfully", zap.Any("comment", newComment))

	return service.toCommentResponse(newComment), nil
}

func (service *commentServiceImpl) Update(c *gin.Context, commentUpdateReq *request.CommentUpdateRequest, userID, commentID uint) (*response.CommentResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var comment entity.Comment

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).First(&comment, "id = ?", commentID).Error; err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok {
				// violation foreign key review_id
				if pgErr.Code == "23503" {
					return exceptions.NewCustomError(http.StatusNotFound, "Review not found")
				}
			}
			return exceptions.NewCustomError(http.StatusNotFound, "Comment not found")
		}

		comment.Content = commentUpdateReq.Content

		if err := tx.Save(&comment).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Comment{}).Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, username")
		}).Take(&comment).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("comment updated successfully", zap.Uint("commentID", commentID))

	return service.toCommentResponse(&comment), nil
}

func (service *commentServiceImpl) Delete(c *gin.Context, userID, commentID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	result := db.Where("id = ?", commentID).Where("user_id = ?", userID).Delete(&entity.Comment{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return exceptions.NewCustomError(http.StatusNotFound, "Comment not found")
	}

	logger.Info("Comment deleted successfully", zap.Uint("commentID", commentID))

	return nil
}

func (service *commentServiceImpl) FindByReviewId(c *gin.Context, reviewID uint) (*[]response.CommentResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var comments []entity.Comment
	if err := db.Model(&entity.Comment{}).
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, username")
		}).
		Where("review_id = ?", reviewID).
		Find(&comments).Error; err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, exceptions.NewCustomError(http.StatusNotFound, "Comments not found")
	}

	logger.Info("Comments fetched successfully", zap.Uint("reviewID", reviewID), zap.Int("count", len(comments)))

	var commentResponses []response.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, *service.toCommentResponse(&comment))
	}

	return &commentResponses, nil
}

func (service *commentServiceImpl) toCommentEntity(req any) *entity.Comment {
	switch v := req.(type) {
	case *request.CommentCreateRequest:
		return &entity.Comment{
			ReviewID: v.ReviewID,
			Content:  v.Content,
		}
	case *request.CommentUpdateRequest:
		return &entity.Comment{
			Content: v.Content,
		}
	default:
		return nil
	}
}

func (service *commentServiceImpl) toCommentResponse(comment *entity.Comment) *response.CommentResponse {
	return &response.CommentResponse{
		ID:       comment.ID,
		ReviewID: comment.ReviewID,
		User: response.CommentUserResponse{
			ID:       comment.User.ID,
			Username: comment.User.Username,
		},
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
