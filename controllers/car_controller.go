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

type CarController interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
}

type carControllerImpl struct {
	services.CarService
}

func NewCarController(carService services.CarService) CarController {
	return &carControllerImpl{carService}
}

// Create car godoc
// @Summary Create car.
// @Description Create a car.
// @Tags Cars
// @Param Body body request.CarCreateRequest true "the body to create a car"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[response.CarResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars [post]
func (controller *carControllerImpl) Create(c *gin.Context) {
	var carCreateReq request.CarCreateRequest

	err := c.ShouldBindJSON(&carCreateReq)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	carRes, err := controller.CarService.Create(c, &carCreateReq)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusCreated, carRes)
}

// Update car godoc
// @Summary Update car.
// @Description Update a car.
// @Tags Cars
// @Param id path int true "Car ID"
// @Param Body body request.CarUpdateRequest true "the body to update a car"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.CarResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars/{id} [patch]
func (controller *carControllerImpl) Update(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	var carUpdateReq request.CarUpdateRequest

	err = c.ShouldBindJSON(&carUpdateReq)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	carRes, err := controller.CarService.Update(c, &carUpdateReq, uint(userID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, carRes)
}

// Delete car godoc
// @Summary Delete car.
// @Description Delete a car.
// @Tags Cars
// @Param id path int true "Car ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars/{id} [delete]
func (controller *carControllerImpl) Delete(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	err = controller.CarService.Delete(c, uint(userID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "car deleted")
}

// Find all car godoc
// @Summary Find all car.
// @Description Find all car.
// @Tags Cars
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.CarResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars [get]
func (controller *carControllerImpl) FindAll(c *gin.Context) {
	cars, err := controller.CarService.FindAll(c)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, cars)
}

// Find car godoc
// @Summary Find car.
// @Description Find a car by id.
// @Tags Cars
// @Param id path int true "Car ID"
// @Produce json
// @Success 200 {object} web.WebSuccess[response.CarResponse]
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars/{id} [get]
func (controller *carControllerImpl) FindById(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	helper.PanicIfError(err)

	car, err := controller.CarService.FindByID(c, uint(userID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, car)
}
