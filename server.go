package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/raihanmd/car-review-sb/app"
	"github.com/raihanmd/car-review-sb/docs"
	"github.com/raihanmd/car-review-sb/helper"
)

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		err := godotenv.Load()
		helper.PanicIfError(err)
	}
}

func main() {
	docs.SwaggerInfo.Title = "Car Review REST API"
	docs.SwaggerInfo.Description = "This is a Car Review REST API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helper.MustGetEnv("SERVER_HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", helper.GetEnv("SERVER_PORT", "3000")))
}
