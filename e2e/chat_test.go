package e2e

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestChat(t *testing.T) {
	must := require.New(t)

	deps := dependencies.Init()

	clearStorage := func() {
		err := deps.DB.ClearTable(context.Background())
		must.Nil(err)
	}

	t.Run("GET /:username/chat-history", func(t *testing.T) {
		defer clearStorage()

		chatMsg := model.ChatMessage{
			MessageId: uuid.New().String(),
			CreatedAt: time.Now().UnixNano(),
			Body:      "Hello, world!",
			UserId:    "esia",
			ConvoId:   uuid.New().String(),
		}
		err := deps.DB.PutItem(context.TODO(), chatMsg)
		must.Nil(err)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/users/esia/chat-history", nil)
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)

		must.Len(chatMessages, 1)
		must.Equal("Hello, world!", chatMessages[0].Body)
		must.Equal("esia", chatMessages[0].UserId)
	})

	t.Run("POST /:username/chat-history", func(t *testing.T) {
		defer clearStorage()

		//chatMsg := model.ChatMessage{
		//	MessageId: uuid.New().String(),
		//	CreatedAt: time.Now().UnixNano(),
		//	Body:      "Hello, world!",
		//	UserId:    "esia",
		//	ConvoId:   uuid.New().String(),
		//}
		//err := deps.DB.PutItem(context.TODO(), chatMsg)
		//must.Nil(err)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/users/esia/chat-history", nil)
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)
		must.Empty(chatMessages)
	})
}
