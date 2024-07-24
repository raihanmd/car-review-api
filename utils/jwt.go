package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
)

var API_SECRET = helper.GetEnv("API_SECRET", "W8j8sLNYNXyhVyjAcyiuaWMCHGFGfcwEG8WsxlOMsPgX0vF73LmSslCaofZls8oNMSmj8bNFnZpxqD3JUUmPhYtRI5gIsSi9riGHTXpgja6RETJiXFI4WTsIfszZcwoW")

func GenerateToken(userId uint, userRole string) (string, error) {
	tokenLifeSpan, err := strconv.Atoi(helper.GetEnv("TOKEN_HOUR_LIFESPAN", "1"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["user_role"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(API_SECRET))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenClaims(c *gin.Context) (id uint, role string, err error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, "", exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", exceptions.NewCustomError(http.StatusBadRequest, "Invalid or expired token")
	}
	userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
	if err != nil {
		return 0, "", err
	}
	return uint(userId), claims["user_role"].(string), nil
}

func UserRoleMustAdmin(c *gin.Context) {
	_, role, err := ExtractTokenClaims(c)
	if err != nil {
		helper.PanicIfError(err)
	}
	if role != entity.RoleAdmin {
		helper.PanicIfError(exceptions.NewCustomError(http.StatusForbidden, "Only admin can manipulate data"))
	}
}
