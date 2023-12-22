package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(username string, signingKey string) (string, error) {
	claims := &JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
