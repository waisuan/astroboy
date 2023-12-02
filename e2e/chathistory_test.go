//go:build e2e
// +build e2e

package e2e

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/webhandlers"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	testify "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestChatHistory(t *testing.T) {
	assert := testify.New(t)
	must := require.New(t)

	t.Run("GET /:username/chat-history", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:username/chat-history")
		c.SetParamNames("username")
		c.SetParamValues("esia")

		deps := dependencies.Init()

		_, err := deps.Db.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(deps.Db.TableName),
			Item: map[string]types.AttributeValue{
				"message_id": &types.AttributeValueMemberS{Value: uuid.New().String()},
				"timestamp":  &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().UnixNano())},
				"body":       &types.AttributeValueMemberS{Value: "Hello, world!"},
				"user_id":    &types.AttributeValueMemberS{Value: uuid.New().String()},
				"convo_id":   &types.AttributeValueMemberS{Value: uuid.New().String()},
			},
		})
		must.Nil(err)

		wh := webhandlers.NewWebHandler(deps)

		assert.NoError(wh.GetChatHistory(c))
		assert.Equal(http.StatusOK, rec.Code)
	})
}
