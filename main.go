package main

import (
	"fmt"

	"github.com/raihanmd/fp-superbootcamp-go/app"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
)

func main() {
	router := app.NewRouter()

	router.Run(fmt.Sprintf(":%s", helper.GetEnv("SERVER_PORT", "3000")))
}
