package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		user_id := claims["user_id"].(float64)
		return user_id
	}

	return 0
}
func ExtractTokenUsername(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		return username
	}

	return ""
}

func CreateToken(user_id uint, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["username"] = username
	claims["expired"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("$p4ssw0rd"))
}
