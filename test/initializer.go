package test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/raihanmd/fp-superbootcamp-go/controllers"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/middlewares"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"github.com/raihanmd/fp-superbootcamp-go/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewUnitTestDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(helper.MustGetEnv("DB_DSN")), &gorm.Config{})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Car{}, &entity.CarSpecification{}, &entity.Brand{}, &entity.Review{}, &entity.Comment{}, &entity.Favourite{}, &entity.Profile{})
	helper.PanicIfError(err)

	db.Exec("CREATE INDEX IF NOT EXISTS idx_title_fulltext ON reviews USING GIN (to_tsvector('english', title))")

	db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")
	db.Exec("CREATE INDEX idx_model_gin ON cars USING GIN (model gin_trgm_ops);")

	return db
}

func NewRouter(db *gorm.DB) *gin.Engine {
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
		OutputPaths: []string{"./log/log.log"},
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

	userService := services.NewUserService()
	carService := services.NewCarService()
	reviewService := services.NewreviewService()
	brandService := services.NewBrandService()
	favouriteService := services.NewFavouriteService()
	commentService := services.NewCommentService()

	// ======================== USER =======================

	userController := controllers.NewUserController(userService, favouriteService, reviewService)

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

	userRouter := apiRouter.Group("/users")

	apiRouter.GET("/users/profile/:id", userController.GetUserProfile)
	apiRouter.GET("/users/favourites", userController.GetFavourites)

	userRouter.Use(middlewares.JwtAuthMiddleware)

	userRouter.PATCH("/password", userController.UpdatePassword)
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

	reviewRouter.GET("/:id/comments", reviewController.FindComments) // comment controller

	reviewRouter.Use(middlewares.JwtAuthMiddleware)

	reviewRouter.POST("/", reviewController.Create)
	reviewRouter.PATCH("/:id", reviewController.Update)
	reviewRouter.DELETE("/:id", reviewController.Delete)

	// ======================== BRAND ROUTE =======================

	brandRouter := apiRouter.Group("/brands")

	brandRouter.GET("/", brandController.FindAll)

	brandRouter.Use(middlewares.JwtAuthMiddleware)

	brandRouter.POST("/", brandController.Create)
	brandRouter.PATCH("/:id", brandController.Update)
	brandRouter.DELETE("/:id", brandController.Delete)

	// ======================== FAVOURITE ROUTE =======================

	favouriteRouter := apiRouter.Group("/favourites")

	favouriteRouter.POST("/:carID", favouriteController.FavouriteCar)
	favouriteRouter.DELETE("/:carID", favouriteController.UnfavouriteCar)

	// ======================== COMMENT ROUTE =======================

	commentRouter := apiRouter.Group("/comments")

	commentRouter.POST("/", commentController.Create)
	commentRouter.PATCH("/:id", commentController.Update)
	commentRouter.DELETE("/:id", commentController.Delete)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return r
}
