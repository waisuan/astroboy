package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	t.Run("generates a valid JWT token", func(t *testing.T) {
		signingKey := "secret"
		tokenString, err := GenerateJwtToken("esia", signingKey)
		require.Nil(t, err)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(signingKey), nil
		})
		require.Nil(t, err)

		require.True(t, token.Valid)

		claims, ok := token.Claims.(jwt.MapClaims)
		require.True(t, ok)
		require.Equal(t, claims["username"].(string), "esia")
		require.NotNil(t, claims["exp"])
	})
}
