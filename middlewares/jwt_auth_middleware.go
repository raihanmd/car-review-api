package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/car-review-sb/model/web"
	"github.com/raihanmd/car-review-sb/utils"
)

func JwtAuthMiddleware(c *gin.Context) {
	err := utils.TokenValid(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &web.WebError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}
	c.Next()
}
