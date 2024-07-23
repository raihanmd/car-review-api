package app

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/raihanmd/fp-superbootcamp-go/controllers"
	"github.com/raihanmd/fp-superbootcamp-go/docs"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/middlewares"
	"github.com/raihanmd/fp-superbootcamp-go/services"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {

	swaggerSchemes := []string{"https"}
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		helper.PanicIfError(err)
		swaggerSchemes = []string{"http"}
	}

	docs.SwaggerInfo.Title = "Car Review REST API"
	docs.SwaggerInfo.Description = "This is a Car Review REST API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helper.MustGetEnv("SERVER_HOST")
	docs.SwaggerInfo.Schemes = swaggerSchemes

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
		OutputPaths: []string{"stdout"},
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

	userService := services.NewUserService()
	carService := services.NewCarService()
	reviewService := services.NewreviewService()
	brandService := services.NewBrandService()
	favouriteService := services.NewFavouriteService()
	commentService := services.NewCommentService()

	// ======================== USER =======================

	userController := controllers.NewUserController(userService, favouriteService)

	// ======================== CARD =======================

	carController := controllers.NewCarController(carService)

	// ======================== REVIEW =======================

	reviewController := controllers.NewreviewController(reviewService, commentService)

	// ======================== BRAND =======================

	brandController := controllers.NewBrandController(brandService)

	// ======================== FAVOURITE =======================

	favouriteController := controllers.NewFavouriteController(favouriteService)

	// ======================== COMMENT =======================

	commentController := controllers.NewCommentController(commentService)

	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "Pragma", "Cache-Control", "Expires"},
			MaxAge:           12 * time.Hour,
		},
	))

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
	apiRouter.POST("/auth/forgot-password", userController.ForgotPassword)
	apiRouter.POST("/auth/reset-password", userController.ResetPassword)

	// ======================== USERS ROUTE =======================

	userRouter := apiRouter.Group("/users")

	apiRouter.GET("/users/profile/:id", userController.GetUserProfile)
	apiRouter.GET("/users/favourites", userController.GetFavourites)
	apiRouter.GET("/users/current", userController.GetCurrentUser)

	userRouter.Use(middlewares.JwtAuthMiddleware)

	userRouter.PATCH("/password", userController.UpdatePassword)
	userRouter.PATCH("/profile", userController.UpdateUserProfile)
	userRouter.DELETE("", userController.DeleteUserProfile)

	// ======================== CARS ROUTE =======================

	carRouter := apiRouter.Group("/cars")

	carRouter.GET("", carController.FindAll)
	carRouter.GET("/:id", carController.FindById)

	carRouter.Use(middlewares.JwtAuthMiddleware)

	carRouter.POST("", carController.Create)
	carRouter.PATCH("/:id", carController.Update)
	carRouter.DELETE("/:id", carController.Delete)

	// ======================== REVIEW ROUTE =======================

	reviewRouter := apiRouter.Group("/reviews")

	reviewRouter.GET("", reviewController.FindAll)
	reviewRouter.GET("/:id", reviewController.FindById)

	reviewRouter.GET("/:id/comments", reviewController.FindComments)
	reviewRouter.Use(middlewares.JwtAuthMiddleware)

	reviewRouter.POST("", reviewController.Create)
	reviewRouter.PATCH("/:id", reviewController.Update)
	reviewRouter.DELETE("/:id", reviewController.Delete)

	// ======================== BRAND ROUTE =======================

	brandRouter := apiRouter.Group("/brands")

	brandRouter.GET("", brandController.FindAll)

	brandRouter.Use(middlewares.JwtAuthMiddleware)

	brandRouter.POST("", brandController.Create)
	brandRouter.PATCH("/:id", brandController.Update)
	brandRouter.DELETE("/:id", brandController.Delete)

	// ======================== FAVOURITE ROUTE =======================

	favouriteRouter := apiRouter.Group("/favourites")

	favouriteRouter.POST("/:carID", favouriteController.FavouriteCar)
	favouriteRouter.DELETE("/:carID", favouriteController.UnfavouriteCar)

	// ======================== COMMENT ROUTE =======================

	commentRouter := apiRouter.Group("/comments")

	commentRouter.POST("", commentController.Create)
	commentRouter.PATCH("/:id", commentController.Update)
	commentRouter.DELETE("/:id", commentController.Delete)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return r
}
