package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/web"
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
	FindComments(*gin.Context)
}

type reviewControllerImpl struct {
	services.ReviewService
	services.CommentService
}

func NewreviewController(reviewService services.ReviewService, commentService services.CommentService) ReviewController {
	return &reviewControllerImpl{reviewService, commentService}
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

	helper.ToResponseJSON(c, http.StatusCreated, reviewRes, nil)
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
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var reviewUpdateReq request.ReviewUpdateRequest

	err = c.ShouldBindJSON(&reviewUpdateReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	reviewRes, err := controller.ReviewService.Update(c, &reviewUpdateReq, userID, uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, reviewRes, nil)
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
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.ReviewService.Delete(c, userID, uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "review deleted", nil)
}

// Find all review godoc
// @Summary Find all review.
// @Description Find all review.
// @Tags Reviews
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Param title query string false "Title"
// @Param car_id query string false "Car ID"
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.FindReviewResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews [get]
func (controller *reviewControllerImpl) FindAll(c *gin.Context) {
	var pagination web.PaginationRequest
	var reviewQueryReq request.ReviewQueryRequest

	if err := c.ShouldBindQuery(&pagination); err != nil {
		panic(err)
	}

	if err := c.ShouldBindQuery(&reviewQueryReq); err != nil {
		panic(err)
	}

	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	reviews, metadata, err := controller.ReviewService.FindAll(c, &reviewQueryReq, &pagination)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, reviews, metadata)
}

// Find review godoc
// @Summary Find review.
// @Description Find a review by id.
// @Tags Reviews
// @Param id path int true "Review ID"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.FindReviewResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews/{id} [get]
func (controller *reviewControllerImpl) FindById(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	review, err := controller.ReviewService.FindByID(c, uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, review, nil)
}

// Find comment by review id godoc
// @Summary Find comment by review id.
// @Description Find a comment by review id.
// @Tags Reviews
// @Param id path int true "Review ID"
// @Produce json
// @Success 201 {object} web.WebSuccess[[]response.CommentResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/reviews/{id}/comments [get]
func (controller *reviewControllerImpl) FindComments(c *gin.Context) {
	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "reviewID must be an integer"))
	}

	comments, err := controller.CommentService.FindByReviewId(c, uint(reviewID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, comments, nil)
}
