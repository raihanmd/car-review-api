package api

import (
	"net/http"

	"github.com/raihanmd/fp-superbootcamp-go/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin := app.NewRouter()

	gin.ServeHTTP(w, r)
}
