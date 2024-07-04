package helper

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetDBAndLogger(c *gin.Context) (*gorm.DB, *zap.Logger) {
	return c.MustGet("db").(*gorm.DB), c.MustGet("logger").(*zap.Logger)
}
