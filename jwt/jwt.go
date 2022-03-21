package jwt

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pytux/twittor/models"
)

func GenerateToken(user models.User) (string, error) {

	key := os.Getenv("tokenKey")

	if key == "" {
		key = "SuperSecretStringToEncodeToken"
	}

	tokenKey := []byte(key)

	payload := jwt.MapClaims{
		"email":         user.Email,
		"name":          user.Name,
		"username":      user.Username,
		"date_of_birth": user.DateOfBirth,
		"bio":           user.Biography,
		"ubication":     user.Ubication,
		"website":       user.WebSite,
		"_id":           user.ID.Hex(),
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(tokenKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
