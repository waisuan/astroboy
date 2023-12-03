//go:build e2e
// +build e2e

package e2e

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"astroboy/internal/webhandlers"
	"context"
	"encoding/json"
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

		chatMsg := model.ChatMessage{
			MessageId: uuid.New().String(),
			CreatedAt: time.Now().UnixNano(),
			Body:      "Hello, world!",
			UserId:    "esia",
			ConvoId:   uuid.New().String(),
		}
		err := deps.Db.PutItem(context.TODO(), chatMsg)
		must.Nil(err)

		wh := webhandlers.NewWebHandler(deps)

		assert.NoError(wh.GetChatHistory(c))
		assert.Equal(http.StatusOK, rec.Code)

		var chatMessages []model.ChatMessage
		_ = json.Unmarshal([]byte(rec.Body.String()), &chatMessages)
		must.Len(chatMessages, 1)

		assert.Equal("Hello, world!", chatMessages[0].Body)
		assert.Equal("esia", chatMessages[0].UserId)
	})
}
