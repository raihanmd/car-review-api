package services

import (
	"html"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"github.com/raihanmd/fp-superbootcamp-go/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	Register(*gin.Context, *entity.User) (*response.RegisterResponse, error)
	Login(*gin.Context, string, string) (*response.LoginResponse, error)
	UpdatePassword(*gin.Context, uint, string) error
	GetUserProfile(*gin.Context, uint) (*response.UserProfileResponse, error)
	UpdateUserProfile(*gin.Context, *entity.User, uint) (*response.UpdateUserProfileResponse, error)
	DeleteUserProfile(*gin.Context, uint) error
}

type userServiceImpl struct{}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (service *userServiceImpl) Register(c *gin.Context, user *entity.User) (*response.RegisterResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = db.Create(user).Error; err != nil {
			return exceptions.NewCustomError(http.StatusConflict, "username already exists")
		}

		if err = db.Create(&entity.Profile{UserID: user.ID}).Error; err != nil {
			return exceptions.NewCustomError(http.StatusConflict, "username already exists")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success registering user", zap.Uint("userID", user.ID))

	registerResponse := response.RegisterResponse{
		Username: user.Username,
		Role:     user.Role,
	}

	return &registerResponse, nil
}

func (service *userServiceImpl) Login(c *gin.Context, username, password string) (*response.LoginResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var err error

	user := entity.User{}

	if err = db.Model(&entity.User{}).Where("username = ?", username).Take(&user).Error; err != nil {
		return nil, exceptions.NewCustomError(http.StatusUnauthorized, "username or password is incorrect")
	}

	err = helper.VerifyPassword(password, user.Password)
	if err != nil {
		logger.Info("user failed login", zap.Any("error", err))
		return nil, exceptions.NewCustomError(http.StatusUnauthorized, "username or password is incorrect")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	loginResponse := response.LoginResponse{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}

	logger.Info("user success logged in", zap.Uint("userID", user.ID))
	return &loginResponse, nil
}

func (service *userServiceImpl) UpdatePassword(c *gin.Context, userID uint, newPassword string) error {
	db, logger := helper.GetDBAndLogger(c)

	hashedPassword, err := helper.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := db.Model(&entity.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	logger.Info("user password updated successfully", zap.Uint("userID", userID))

	return nil
}

func (service *userServiceImpl) GetUserProfile(c *gin.Context, userID uint) (*response.UserProfileResponse, error) {
	db, _ := helper.GetDBAndLogger(c)

	var responseUser response.UserProfileResponse

	if err := db.Model(&entity.User{}).
		Select("users.id, users.username, users.role, profiles.user_id, profiles.email, profiles.full_name, profiles.bio, profiles.age, profiles.gender").
		Joins("left join profiles on users.id = profiles.user_id").
		Where("users.id = ?", userID).
		Scan(&responseUser).Error; err != nil {
		return nil, err
	}

	if responseUser.ID == 0 {
		return nil, exceptions.NewCustomError(http.StatusNotFound, "user not found")
	}

	return &responseUser, nil
}

func (service *userServiceImpl) UpdateUserProfile(c *gin.Context, updatedUser *entity.User, userID uint) (*response.UpdateUserProfileResponse, error) {
	db, logger := helper.GetDBAndLogger(c)

	var responseUser response.UpdateUserProfileResponse

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.User{ID: userID}).
			Omit("password").
			Updates(updatedUser).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Profile{}).
			Where("user_id = ?", userID).
			Updates(updatedUser.Profile).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.User{}).
			Select("profiles.*, users.id, users.username, users.role").
			Joins("left join profiles on users.id = profiles.user_id").
			Take(&responseUser, "users.id = ?", userID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	logger.Info("success updating user profile", zap.Uint("userID", userID))

	return &responseUser, nil
}

func (service *userServiceImpl) DeleteUserProfile(c *gin.Context, userID uint) error {
	db, logger := helper.GetDBAndLogger(c)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&entity.Profile{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&entity.User{ID: userID}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	logger.Info("user profile and user deleted successfully", zap.Uint("userID", userID))
	return nil
}
