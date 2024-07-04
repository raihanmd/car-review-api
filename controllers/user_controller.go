package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
	_ "github.com/raihanmd/car-review-sb/model/web"
	"github.com/raihanmd/car-review-sb/model/web/request"
	_ "github.com/raihanmd/car-review-sb/model/web/response"
	"github.com/raihanmd/car-review-sb/services"
	"github.com/raihanmd/car-review-sb/utils"
)

type UserController interface {
	Register(*gin.Context)
	Login(c *gin.Context)
	UpdatePassword(c *gin.Context)
	GetUserProfile(c *gin.Context)
	UpdateUserProfile(c *gin.Context)
	DeleteUserProfile(c *gin.Context)
}

type userControllerImpl struct {
	services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userControllerImpl{
		userService,
	}
}

// Register godoc
// @Summary User Register.
// @Description Registering a user from public access.
// @Tags Auth
// @Param Body body request.RegisterRequest true "the body to register a user"
// @Produce json
// @Success 201 {object} web.WebSuccess[response.RegisterResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Router /api/auth/register [post]
func (controller *userControllerImpl) Register(c *gin.Context) {
	var registerReq request.RegisterRequest

	err := c.ShouldBindJSON(&registerReq)
	helper.PanicIfError(err)

	newUser := entity.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
	}

	registerRes, err := controller.UserService.Register(c, &newUser)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusCreated, registerRes)
}

// LoginUser godoc
// @Summary User Login.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body request.LoginRequest true "the body to login a user"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.LoginResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 401 {object} web.WebUnauthorizedError
// @Router /api/auth/login [post]
func (controller *userControllerImpl) Login(c *gin.Context) {
	var loginReq request.LoginRequest

	err := c.ShouldBindJSON(&loginReq)
	helper.PanicIfError(err)

	loginRes, err := controller.UserService.Login(c, loginReq.Username, loginReq.Password)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, loginRes)
}

// UpdatePassword godoc
// @Summary Update user password.
// @Description Update the current user's password.
// @Tags Users
// @Param Body body request.UpdatePasswordRequest true "the body to update a password"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Router /api/users/password [patch]
func (controller *userControllerImpl) UpdatePassword(c *gin.Context) {
	var updatePasswordReq request.UpdatePasswordRequest

	err := c.ShouldBindJSON(&updatePasswordReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.UserService.UpdatePassword(c, userID, updatePasswordReq.Password)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "password updated")
}

// UpdatePassword godoc
// @Summary Get user profile.
// @Description Get user profile data.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} web.WebSuccess[response.UserProfileResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Router /api/users/profile/{id} [get]
func (controller *userControllerImpl) GetUserProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	user, err := controller.UserService.GetUserProfile(c, uint(userID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, user)
}

// UpdateUserProfile godoc
// @Summary Update User Profile.
// @Description Update the profile of a user.
// @Tags Users
// @Param Body body request.UpdateUserProfileRequest true "the body to update user profile"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.UpdateUserProfileResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Router /api/users/profile [patch]
func (controller *userControllerImpl) UpdateUserProfile(c *gin.Context) {
	var updateUserReq request.UpdateUserProfileRequest

	err := c.ShouldBindJSON(&updateUserReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	updatedProfile := &entity.User{
		Username: updateUserReq.Username,
		Profile: entity.Profile{
			Email:    updateUserReq.Email,
			FullName: updateUserReq.FullName,
			Bio:      updateUserReq.Bio,
			Age:      updateUserReq.Age,
			Gender:   updateUserReq.Gender,
		},

		ID: uint(userID),
	}

	userResponse, err := controller.UserService.UpdateUserProfile(c, updatedProfile, userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, userResponse)
}

// DeleteUserProfile godoc
// @Summary Delete User.
// @Description Delete a user profile by ID.
// @Tags Users
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Router /api/users [delete]
func (controller *userControllerImpl) DeleteUserProfile(c *gin.Context) {
	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.UserService.DeleteUserProfile(c, userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "user deleted")
}
