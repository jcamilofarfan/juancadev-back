package jwt

import (
	"errors"
	"fmt"

	"time"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/config"
	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	ID uint
}

func Generate(payload *TokenPayload) (string, error) {
	v, err := time.ParseDuration(config.TOKENEXP)

	if (payload.ID == 0){
		return "", fmt.Errorf("Payload id cannot be 0")
	}

	if err != nil {
		return "", fmt.Errorf("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(config.TOKENKEY))

	if err != nil {
		return "", fmt.Errorf("%v",err)
	}

	return token, nil
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.TOKENKEY), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("Something went wrong")
	}

	return &TokenPayload{
		ID: uint(id),
	}, nil
}
