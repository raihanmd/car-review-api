package exceptions

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raihanmd/car-review-sb/model/web"
	"go.uber.org/zap"
)

func GlobalErrorHandler(c *gin.Context) {
	logger := c.MustGet("logger").(*zap.Logger)

	defer func() {
		if err := recover(); err != nil {
			logger.Error("global error handler", zap.Any("error", err))
			switch e := err.(type) {
			case *customError:
				c.AbortWithStatusJSON(e.Code, &web.WebError{
					Code:   e.Code,
					Errors: e.Errors,
				})
			case validator.ValidationErrors:
				HandleValidationErrors(c, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, &web.WebError{
					Code:   http.StatusInternalServerError,
					Errors: http.StatusText(http.StatusInternalServerError),
				})
			}
		}
	}()

	c.Next()
}

func HandleValidationErrors(c *gin.Context, err validator.ValidationErrors) {
	errors := make(map[string]string)
	for _, fe := range err {
		field := toSnakeCase(fe.Field())
		errors[field] = "Field validation for '" + field + "' failed on the '" + fe.Tag() + "' tag"
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, &web.WebError{
		Code:   http.StatusBadRequest,
		Errors: errors,
	})
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, rune := range s {
		if i > 0 && i < len(s)-1 && 'A' <= rune && rune <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(rune)
	}
	return strings.ToLower(result.String())
}
