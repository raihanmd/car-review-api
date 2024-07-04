package test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/raihanmd/car-review-sb/app"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/entity"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)
	m.Run()
}

func TestDatabase(t *testing.T) {
	db := app.NewConnection()

	var user entity.User

	db.Select("email").Where("id = ?", 1).Find(&user)

	t.Log(user)
}
