package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/fp-superbootcamp-go/model/web"
)

func ToResponseJSON[T any](c *gin.Context, code int, data T, pagination *web.Metadata) {
	if pagination == nil {
		c.JSON(code, web.WebSuccess[T]{
			Code:     code,
			Message:  http.StatusText(code),
			Data:     data,
			Metadata: nil,
		})

		return
	}

	c.JSON(code, web.WebSuccess[T]{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
		Metadata: &web.Metadata{
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalPages: pagination.TotalPages,
			TotalData:  pagination.TotalData,
		},
	})
}
