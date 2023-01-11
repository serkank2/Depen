package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now()
	tokenString, err := token.SignedString([]byte("Kaplan.42"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
