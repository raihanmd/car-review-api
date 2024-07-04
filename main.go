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

}

func main() {
	swaggerSchemes := []string{"htpps"}
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		helper.PanicIfError(err)
		swaggerSchemes = []string{"http", "https"}
	}

	docs.SwaggerInfo.Title = "Car Review REST API"
	docs.SwaggerInfo.Description = "This is a Car Review REST API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = helper.MustGetEnv("SERVER_HOST")
	docs.SwaggerInfo.Schemes = swaggerSchemes

	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", helper.GetEnv("SERVER_PORT", "3000")))
}
