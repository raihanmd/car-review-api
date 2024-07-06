package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	_ "github.com/raihanmd/car-review-sb/model/web"
	"github.com/raihanmd/car-review-sb/services"
	"github.com/raihanmd/car-review-sb/utils"
)

type FavouriteController interface {
	FavouriteCar(c *gin.Context)
	UnfavouriteCar(c *gin.Context)
}

type favouriteControllerImpl struct {
	services.FavouriteService
}

func NewFavouriteController(favouriteService services.FavouriteService) FavouriteController {
	return &favouriteControllerImpl{favouriteService}
}

// Favourite a car godoc
// @Summary Favourite a car.
// @Description Favourite a car.
// @Tags Favourites
// @Param carID path int true "car ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/favourites/{carID} [post]
func (controller *favouriteControllerImpl) FavouriteCar(c *gin.Context) {
	carID, err := strconv.ParseUint(c.Param("carID"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "carID must be an integer"))
	}

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.FavouriteService.FavouriteCar(c, uint(carID), userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, fmt.Sprintf("success add favourite for carID %v by userID %v", carID, userID))
}

// Unfavourite a car godoc
// @Summary Unfavourite a car.
// @Description Unfavourite a car.
// @Tags Favourites
// @Param carID path int true "car ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 403 {object} web.WebForbiddenError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/favourites/{carID} [delete]
func (controller *favouriteControllerImpl) UnfavouriteCar(c *gin.Context) {
	carID, err := strconv.ParseUint(c.Param("carID"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "carID must be an integer"))
	}

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.FavouriteService.UnfavouriteCar(c, uint(carID), userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, fmt.Sprintf("success unfavourite for carID %v by userID %v", carID, userID))
}
