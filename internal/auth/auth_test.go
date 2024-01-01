package auth

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/mocks"
	"astroboy/internal/model"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthService_RegisterUser(t *testing.T) {
	t.Run("adds new user to the DB", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().PutItem(gomock.Any(), gomock.AssignableToTypeOf(&model.User{}), gomock.Any()).Return(nil)

		as := NewAuthService(&dependencies.Dependencies{DB: mockDb})
		out := as.RegisterUser("esia", "password", "testing@mail.com")
		require.Nil(t, out)
	})
}

func TestAuthService_LoginUser(t *testing.T) {
	t.Run("valid password", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(dependencies.DbQueryOutput{
			{
				"id":        &types.AttributeValueMemberS{Value: "esia"},
				"timestamp": &types.AttributeValueMemberN{Value: "0"},
				"password":  &types.AttributeValueMemberS{Value: "$2a$14$MXE5eZZeD1BzA6tiIyB0/.SpMj5lDFFL4fFHpcTKvfLj9MWk89eVe"},
			},
		}, nil)

		as := NewAuthService(&dependencies.Dependencies{DB: mockDb})

		out := as.LoginUser("esia", "password")
		require.Nil(t, out)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(dependencies.DbQueryOutput{
			{
				"id":        &types.AttributeValueMemberS{Value: "esia"},
				"timestamp": &types.AttributeValueMemberN{Value: "0"},
				"password":  &types.AttributeValueMemberS{Value: "$2a$14$MXE5eZZeD1BzA6tiIyB0/.SpMj5lDFFL4fFHpcTKvfLj9MWk89eVe"},
			},
		}, nil)

		as := NewAuthService(&dependencies.Dependencies{DB: mockDb})

		out := as.LoginUser("esia", "dummy")
		require.NotNil(t, out)
	})

	t.Run("user does not exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(dependencies.DbQueryOutput{}, nil)

		as := NewAuthService(&dependencies.Dependencies{DB: mockDb})

		out := as.LoginUser("esia", "dummy")
		require.NotNil(t, out)
		require.Equal(t, "unable to locate user record in DB", out.Error())
	})
}

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
