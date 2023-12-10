package e2e

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
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

	t.Run("GET /:username/chat-history - has chat messages", func(t *testing.T) {
		defer clearStorage()

		userId := "esia"
		convoId := uuid.New().String()
		numOfMessages := 5

		for i := 0; i < numOfMessages; i++ {
			chatMsg := model.ChatMessage{
				MessageId: uuid.New().String(),
				CreatedAt: time.Now().UnixNano(),
				Body:      gofakeit.RandomString([]string{"Hello, world!", "Lorem Ipsum", "How are you?"}),
				UserId:    userId,
				ConvoId:   convoId,
			}
			err := deps.DB.PutItem(context.TODO(), chatMsg)
			must.Nil(err)
		}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-history", userId), nil)
		must.Nil(err)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)

		must.Len(chatMessages, numOfMessages)

		for _, m := range chatMessages {
			must.NotEmpty(m.Body)
			must.NotEmpty(m.MessageId)
			must.NotEmpty(m.CreatedAt)
			must.NotEmpty(m.ConvoId)
			must.Equal(userId, m.UserId)
		}
	})

	t.Run("GET /:username/chat-history - has no chat messages", func(t *testing.T) {
		defer clearStorage()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-history", "rando"), nil)
		must.Nil(err)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)

		must.Empty(chatMessages)
	})

	t.Run("POST /:username/chat-history - successful", func(t *testing.T) {
		defer clearStorage()

		body := []byte(`{
			"body": "Hello, world!"
		}`)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/users/esia/chat-message", bytes.NewBuffer(body))
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusCreated, res.StatusCode)
	})

	t.Run("POST /:username/chat-history - malformed request", func(t *testing.T) {
		defer clearStorage()

		body := []byte(`Hello, world!`)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/users/esia/chat-message", bytes.NewBuffer(body))
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusInternalServerError, res.StatusCode)
	})
}
