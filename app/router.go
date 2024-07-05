package app

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/raihanmd/car-review-sb/controllers"
	"github.com/raihanmd/car-review-sb/exceptions"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/middlewares"
	"github.com/raihanmd/car-review-sb/services"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("no_space", func(fl validator.FieldLevel) bool {
			return !strings.Contains(fl.Field().String(), " ")
		})
		v.RegisterValidation("lowercase", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == strings.ToLower(fl.Field().String())
		})
		v.RegisterValidation("uppercase", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == strings.ToUpper(fl.Field().String())
		})
		v.RegisterValidation("url", func(fl validator.FieldLevel) bool {
			_, err := url.ParseRequestURI(fl.Field().String())
			return err == nil
		})
	}

	cfg := zap.Config{
		OutputPaths: []string{"./log/test.log", "stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "time_stamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
	}

	logger, err := cfg.Build()
	helper.PanicIfError(err)
	defer logger.Sync()

	db := NewConnection()

	// ======================== USER =======================

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// ======================== CARD =======================

	carService := services.NewCarService()
	carController := controllers.NewCarController(carService)

	// ======================== REVIEW =======================

	reviewService := services.NewreviewService()
	reviewController := controllers.NewreviewController(reviewService)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("logger", logger)
	})

	r.Use(exceptions.GlobalErrorHandler)

	r.NoRoute(func(c *gin.Context) {
		panic(exceptions.NewCustomError(http.StatusNotFound, fmt.Sprintf("path not found, use https://%s for API docs", helper.MustGetEnv("SERVER_HOST")+"/docs/index.html")))
	})

	apiRouter := r.Group("/api")

	// ======================== AUTH ROUTE =======================

	apiRouter.POST("/auth/register", userController.Register)
	apiRouter.POST("/auth/login", userController.Login)

	// ======================== USERS ROUTE =======================

	apiRouter.GET("/users/profile/:id", userController.GetUserProfile)

	userRouter := apiRouter.Group("/users", middlewares.JwtAuthMiddleware)

	userRouter.PUT("/password", userController.UpdatePassword)
	userRouter.PATCH("/profile", userController.UpdateUserProfile)
	userRouter.DELETE("/", userController.DeleteUserProfile)

	// ======================== CARS ROUTE =======================

	carRouter := apiRouter.Group("/cars")

	carRouter.GET("/", carController.FindAll)
	carRouter.GET("/:id", carController.FindById)

	carRouter.Use(middlewares.JwtAuthMiddleware)

	carRouter.POST("/", carController.Create)
	carRouter.PATCH("/:id", carController.Update)
	carRouter.DELETE("/:id", carController.Delete)

	// ======================== REVIEW ROUTE =======================

	reviewRouter := apiRouter.Group("/reviews")

	reviewRouter.GET("/", reviewController.FindAll)
	reviewRouter.GET("/:id", reviewController.FindById)

	reviewRouter.Use(middlewares.JwtAuthMiddleware)

	reviewRouter.POST("/", reviewController.Create)
	reviewRouter.PATCH("/:id", reviewController.Update)
	reviewRouter.DELETE("/:id", reviewController.Delete)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return r
}
