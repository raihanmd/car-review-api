package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	_ "github.com/raihanmd/car-review-sb/model/web"
	"github.com/raihanmd/car-review-sb/model/web/request"
	_ "github.com/raihanmd/car-review-sb/model/web/response"
	"github.com/raihanmd/car-review-sb/services"
	"github.com/raihanmd/car-review-sb/utils"
)

type BrandController interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	FindAll(*gin.Context)
}

type brandControllerImpl struct {
	services.BrandService
}

func NewBrandController(brandService services.BrandService) BrandController {
	return &brandControllerImpl{brandService}
}

// Create brand godoc
// @Summary Create brand.
// @Description Create a brand.
// @Tags Brands
// @Param Body body request.BrandRequest true "the body to create a brand"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[response.BrandResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/brands [post]
func (controller *brandControllerImpl) Create(c *gin.Context) {
	var brandCreateReq request.BrandRequest

	err := c.ShouldBindJSON(&brandCreateReq)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	brandRes, err := controller.BrandService.Create(c, &brandCreateReq)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusCreated, brandRes, nil)
}

// Update brand godoc
// @Summary Update brand.
// @Description Update a brand.
// @Tags Brands
// @Param id path int true "brand ID"
// @Param Body body request.BrandRequest true "the body to update a brand"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[response.BrandResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/brands/{id} [patch]
func (controller *brandControllerImpl) Update(c *gin.Context) {
	brandID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var brandUpdateReq request.BrandRequest

	err = c.ShouldBindJSON(&brandUpdateReq)
	helper.PanicIfError(err)

	utils.UserRoleMustAdmin(c)

	brandRes, err := controller.BrandService.Update(c, &brandUpdateReq, uint(brandID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, brandRes, nil)
}

// Delete brand godoc
// @Summary Delete brand.
// @Description Delete a brand.
// @Tags Brands
// @Param id path int true "brand ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/brands/{id} [delete]
func (controller *brandControllerImpl) Delete(c *gin.Context) {
	brandID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	utils.UserRoleMustAdmin(c)

	err = controller.BrandService.Delete(c, uint(brandID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "brand deleted", nil)
}

// Find all brand godoc
// @Summary Find all brand.
// @Description Find all brand.
// @Tags Brands
// @Produce json
// @Success 200 {object} web.WebSuccess[[]response.BrandResponse]
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/brands [get]
func (controller *brandControllerImpl) FindAll(c *gin.Context) {
	brands, err := controller.BrandService.FindAll(c)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, brands, nil)
}
