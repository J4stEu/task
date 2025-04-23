package session

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func (a *Session) GetTokenValue(r *http.Request) (string, error) {
	tokenValue, err := a.extractTokenValue(r)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errExtractToken, err)
	}

	return tokenValue, nil
}

func (a *Session) extractTokenValue(r *http.Request) (string, error) {
	tokenCookie, err := a.GetCookie(r)
	if err != nil {
		return "", err
	}

	tokenString := tokenCookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.apiSecret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		value := fmt.Sprintf("%v", claims["token"])
		return value, nil
	}

	return "", errInvalidToken
}
