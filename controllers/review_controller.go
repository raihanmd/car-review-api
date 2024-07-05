package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/helper"
	_ "github.com/raihanmd/car-review-sb/model/web"
	"github.com/raihanmd/car-review-sb/model/web/request"
	_ "github.com/raihanmd/car-review-sb/model/web/response"
	"github.com/raihanmd/car-review-sb/services"
	"github.com/raihanmd/car-review-sb/utils"
)

type ReviewController interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
}

type reviewControllerImpl struct {
	services.ReviewService
}

func NewreviewController(reviewService services.ReviewService) ReviewController {
	return &reviewControllerImpl{reviewService}
}

// Create review godoc
// @Summary Create review.
// @Description Create a review.
// @Tags Reviews
// @Param Body body request.ReviewCreateRequest true "the body to create a review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[response.ReviewResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews [post]
func (controller *reviewControllerImpl) Create(c *gin.Context) {
	var reviewCreateReq request.ReviewCreateRequest

	err := c.ShouldBindJSON(&reviewCreateReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	reviewRes, err := controller.ReviewService.Create(c, &reviewCreateReq, userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusCreated, reviewRes)
}

// Update review godoc
// @Summary Update review.
// @Description Update a review.
// @Tags Reviews
// @Param id path int true "Review ID"
// @Param Body body request.ReviewUpdateRequest true "the body to update a review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.ReviewResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews/{id} [patch]
func (controller *reviewControllerImpl) Update(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	var reviewUpdateReq request.ReviewUpdateRequest

	err = c.ShouldBindJSON(&reviewUpdateReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	reviewRes, err := controller.ReviewService.Update(c, &reviewUpdateReq, uint(userID), uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, reviewRes)
}

// Delete review godoc
// @Summary Delete review.
// @Description Delete a review.
// @Tags Reviews
// @Param id path int true "Review ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews/{id} [delete]
func (controller *reviewControllerImpl) Delete(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.ReviewService.Delete(c, uint(userID), uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "review deleted")
}

// Find all review godoc
// @Summary Find all review.
// @Description Find all review.
// @Tags Reviews
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.FindReviewResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews [get]
func (controller *reviewControllerImpl) FindAll(c *gin.Context) {
	reviews, err := controller.ReviewService.FindAll(c)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, reviews)
}

// Find review godoc
// @Summary Find review.
// @Description Find a review by id.
// @Tags Reviews
// @Param id path int true "review ID"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.FindReviewResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews/{id} [get]
func (controller *reviewControllerImpl) FindById(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	review, err := controller.ReviewService.FindByID(c, uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, review)
}
