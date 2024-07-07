package httptools

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func GetValueFromJwtToken(r *http.Request, jwtSecretKey []byte, field string) (string, error) {
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		return "", fmt.Errorf("cannot get access token: %v", err)
	}
	token, _ := jwt.Parse(accessToken.Value, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unable to retrieve claims from JWT token")
	}

	username, ok := claims[field].(string)
	if !ok {
		return "", fmt.Errorf("unable to retrieve %s from JWT token", field)
	}
	return username, nil
}
