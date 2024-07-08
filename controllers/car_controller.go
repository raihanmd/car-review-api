package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/web"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/request"
	_ "github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"github.com/raihanmd/fp-superbootcamp-go/services"
	"github.com/raihanmd/fp-superbootcamp-go/utils"
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

	helper.ToResponseJSON(c, http.StatusCreated, carRes, nil)
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
	carID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var carUpdateReq request.CarUpdateRequest

	err = c.ShouldBindJSON(&carUpdateReq)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	carRes, err := controller.CarService.Update(c, &carUpdateReq, uint(carID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, carRes, nil)
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
	carID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	utils.UserRoleMustAdmin(c)

	err = controller.CarService.Delete(c, uint(carID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "car deleted", nil)
}

// Find all car godoc
// @Summary Find all car.
// @Description Find all car.
// @Tags Cars
// @Param limit query int false "Limit" default(10)
// @Param page query int false "Page" default(1)
// @Param brand_id query int false "Brand ID"
// @Param model query string false "Model"
// @Param min_year query int false "Minimum Year"
// @Param max_year query int false "Maximum Year"
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.CarResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/cars [get]
func (controller *carControllerImpl) FindAll(c *gin.Context) {
	var pagination web.PaginationRequest
	var carQueryReq request.CarQueryRequest

	if err := c.ShouldBindQuery(&pagination); err != nil {
		panic(err)
	}

	if err := c.ShouldBindQuery(&carQueryReq); err != nil {
		panic(err)
	}

	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	cars, metadata, err := controller.CarService.FindAll(c, &carQueryReq, &pagination)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, cars, metadata)
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
	carID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	car, err := controller.CarService.FindByID(c, uint(carID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, car, nil)
}
