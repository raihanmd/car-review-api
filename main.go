package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/raihanmd/fp-superbootcamp-go/app"
	"github.com/raihanmd/fp-superbootcamp-go/docs"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
)

func init() {

}

func main() {
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

	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", helper.GetEnv("SERVER_PORT", "3000")))
}
