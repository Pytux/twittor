package routers

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/models"
)

var (
	Email  string
	UserID string
)

func ProcessToken(token string) (*models.Claim, bool, string, error) {

	key := os.Getenv("tokenKey")

	if key == "" {
		key = "SuperSecretStringToEncodeToken"
	}

	tokenKey := []byte(key)

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err == nil {
		_, exist, _ := db.CheckIfEmailIsInUse(claims.Email)
		if exist {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, exist, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
