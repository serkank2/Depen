package auth

import "github.com/golang-jwt/jwt"

func ParseToken(token string) (jwt.MapClaims, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("Kaplan.42"), nil
	})
	if err != nil {
		return nil, err
	}
	calms := parse.Claims.(jwt.MapClaims)
	return calms, nil
}
