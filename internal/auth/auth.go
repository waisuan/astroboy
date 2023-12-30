package auth

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	deps *dependencies.Dependencies
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthService(deps *dependencies.Dependencies) *AuthService {
	return &AuthService{deps: deps}
}

func (as *AuthService) RegisterUser(username string, password string, email string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u := &model.User{
		Id:           username,
		Timestamp:    0,
		Password:     string(hash),
		Email:        email,
		RegisteredAt: time.Now().UnixNano(),
	}

	cond := expression.AttributeNotExists(expression.Name(dependencies.PartitionKey))
	expr, err := expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		return err
	}

	return as.deps.DB.PutItem(context.TODO(), u, &expr)
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
