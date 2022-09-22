package middleware

import (
	"github.com/golang-jwt/jwt/v4"
)

type cliams struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func VerifyToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("screte"), nil
	})

	if err != nil {
		return t, err
	}
	
	return t, nil
}

func CreateToken(id string) (interface{}, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, cliams{id, jwt.RegisteredClaims{}})
	token, err := t.SignedString([]byte("screte"))

	if err != nil {
		return nil, err
	}

	return token, nil
}
