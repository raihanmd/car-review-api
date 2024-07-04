package services

import (
	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/model/entity"
	"github.com/raihanmd/car-review-sb/model/web/response"
)

type CarService interface {
	Create(*gin.Context, *entity.Car) (*response.CarResponse, error)
}
