package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/raihanmd/car-review-sb/helper"
	"github.com/raihanmd/car-review-sb/model/web/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Router *gin.Engine

var AuthorizedToken string

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	helper.PanicIfError(err)

	DB = NewUnitTestDatabase()
	Router = NewRouter(DB)

	TruncateUser(DB)
	CreateRootUser(DB)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("should success register", func(t *testing.T) {
		newUser := request.RegisterRequest{
			Username: "test",
			Password: "testtest",
		}
		requestBody, _ := json.Marshal(newUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 201, response.StatusCode)
		assert.Equal(t, 201.0, jsonResult["code"])
		assert.Equal(t, "Created", jsonResult["message"])
		assert.Equal(t, newUser.Username, jsonResult["data"].(map[string]any)["username"])
		assert.Equal(t, "USER", jsonResult["data"].(map[string]any)["role"])
	})

	t.Run("should error if password is less than 6 or username less than 3", func(t *testing.T) {
		newUser := request.RegisterRequest{
			Username: "ra",
			Password: "test",
		}
		requestBody, _ := json.Marshal(newUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})

	t.Run("should error if username already exists", func(t *testing.T) {
		newUser := request.RegisterRequest{
			Username: "test",
			Password: "testtest",
		}
		requestBody, _ := json.Marshal(newUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 409, response.StatusCode)
		assert.Equal(t, 409.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})
}

func TestLogin(t *testing.T) {
	t.Run("should success login", func(t *testing.T) {
		loginUser := request.LoginRequest{
			Username: "test",
			Password: "testtest",
		}
		requestBody, _ := json.Marshal(loginUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.Equal(t, "USER", jsonResult["data"].(map[string]any)["role"])
		assert.NotNil(t, jsonResult["data"].(map[string]any)["token"])

		AuthorizedToken = jsonResult["data"].(map[string]any)["token"].(string)
	})

	t.Run("should error if password or username are empty", func(t *testing.T) {
		loginUser := request.LoginRequest{
			Username: "",
			Password: "",
		}
		requestBody, _ := json.Marshal(loginUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})

	t.Run("should error if username or pasword wrong", func(t *testing.T) {
		loginUser := request.RegisterRequest{
			Username: "test",
			Password: "wrong",
		}
		requestBody, _ := json.Marshal(loginUser)

		request := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, 401.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})
}

func TestUpdatePasswordUser(t *testing.T) {
	t.Run("should success update password", func(t *testing.T) {
		newPassword := request.UpdatePasswordRequest{
			Password: "mynewpassword",
		}
		requestBody, _ := json.Marshal(newPassword)

		requestHttp := httptest.NewRequest(http.MethodPatch, "/api/users/password", strings.NewReader(string(requestBody)))
		requestHttp.Header.Add("Content-Type", "application/json")

		requestHttp.Header.Add("Authorization", "Bearer "+AuthorizedToken)

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, requestHttp)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.NotNil(t, jsonResult["data"])

		t.Run("should success login again with updated password", func(t *testing.T) {
			newLogin := request.LoginRequest{
				Username: "test",
				Password: "mynewpassword",
			}

			requestBody, _ = json.Marshal(newLogin)

			request := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(requestBody)))
			request.Header.Add("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			Router.ServeHTTP(recorder, request)

			response := recorder.Result()

			var jsonResult map[string]any

			json.NewDecoder(response.Body).Decode(&jsonResult)

			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, 200.0, jsonResult["code"])
			assert.Equal(t, "OK", jsonResult["message"])
			assert.Equal(t, "USER", jsonResult["data"].(map[string]any)["role"])
			assert.NotNil(t, jsonResult["data"].(map[string]any)["token"])

			AuthorizedToken = jsonResult["data"].(map[string]any)["token"].(string)
		})
	})
}

func TestUpdateUserProfile(t *testing.T) {
	t.Run("should success update profile", func(t *testing.T) {
		var (
			email    = "email@email.com"
			fullName = "My Full Name"
			age      = 20
			gender   = "MALE"
		)

		updateProfile := request.UpdateUserProfileRequest{
			Username: "newusername",
			Email:    &email,
			FullName: &fullName,
			Bio:      nil,
			Age:      &age,
			Gender:   &gender,
		}
		requestBody, _ := json.Marshal(updateProfile)

		requestHttp := httptest.NewRequest(http.MethodPatch, "/api/users/profile", strings.NewReader(string(requestBody)))
		requestHttp.Header.Add("Content-Type", "application/json")

		requestHttp.Header.Add("Authorization", "Bearer "+AuthorizedToken)

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, requestHttp)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.Equal(t, updateProfile.Username, jsonResult["data"].(map[string]any)["username"])
		assert.Equal(t, "USER", jsonResult["data"].(map[string]any)["role"])
		assert.Equal(t, email, jsonResult["data"].(map[string]any)["email"])
		assert.Equal(t, fullName, jsonResult["data"].(map[string]any)["full_name"])
		assert.Equal(t, nil, jsonResult["data"].(map[string]any)["bio"])
		assert.Equal(t, float64(age), jsonResult["data"].(map[string]any)["age"])
		assert.Equal(t, gender, jsonResult["data"].(map[string]any)["gender"])
	})

	t.Run("should error if request is wrong", func(t *testing.T) {
		var (
			email    = "wrong"
			fullName = "My Full Name"
			bio      = "Hello this is my bio"
			age      = 20
			gender   = "wrong"
		)

		updateProfile := request.UpdateUserProfileRequest{
			Username: "newusername",
			Email:    &email,
			FullName: &fullName,
			Bio:      &bio,
			Age:      &age,
			Gender:   &gender,
		}
		requestBody, _ := json.Marshal(updateProfile)

		request := httptest.NewRequest(http.MethodPatch, "/api/users/profile", strings.NewReader(string(requestBody)))
		request.Header.Add("Content-Type", "application/json")

		request.Header.Add("Authorization", "Bearer "+AuthorizedToken)

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})
}

func TestGetProfileById(t *testing.T) {
	t.Run("should success get profile by id", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/users/profile/1", nil)
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 200.0, jsonResult["code"])
		assert.Equal(t, "OK", jsonResult["message"])
		assert.NotNil(t, jsonResult["data"])
	})

	t.Run("should 404 if user not found", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/users/profile/404", nil)
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, 404.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})

	t.Run("should error if param is not int", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/api/users/profile/wrong", nil)
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		Router.ServeHTTP(recorder, request)

		response := recorder.Result()

		var jsonResult map[string]any

		json.NewDecoder(response.Body).Decode(&jsonResult)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400.0, jsonResult["code"])
		assert.NotNil(t, jsonResult["errors"])
	})
}
