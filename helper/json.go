package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/model/web"
)

func ToResponseJSON[T any](c *gin.Context, code int, data T) {
	c.JSON(code, web.WebSuccess[T]{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	})
}
