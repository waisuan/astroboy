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
	"os"
	"testing"
	"time"
)

func TestChat(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	must := require.New(t)
	deps := dependencies.Init()
	userId := "esia"

	clearStorage := func() {
		err := deps.DB.ClearTable(context.Background())
		must.Nil(err)

		os.Unsetenv("APP_ENV")
	}

	login := func(username string) string {
		body := []byte(fmt.Sprintf(`{
			"username": "%s",
            "password": "dummy"
		}`, username))
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:5000/login", bytes.NewBuffer(body))
		res, _ := http.DefaultClient.Do(req)

		token := make(map[string]interface{})
		_ = json.NewDecoder(res.Body).Decode(&token)

		return token["token"].(string)
	}

	t.Run("GET /:username/chat-history - has chat messages", func(t *testing.T) {
		defer clearStorage()

		convoId := uuid.New().String()
		numOfMessages := 5

		authToken := login(userId)

		for i := 0; i < numOfMessages; i++ {
			chatMsg := model.ChatMessage{
				Id:        uuid.New().String(),
				Timestamp: time.Now().UnixNano(),
				Body:      gofakeit.RandomString([]string{"Hello, world!", "Lorem Ipsum", "How are you?"}),
				UserId:    userId,
				ConvoId:   convoId,
			}
			err := deps.DB.PutItem(context.TODO(), chatMsg, nil)
			must.Nil(err)
		}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-history", userId), nil)
		must.Nil(err)

		req.Header.Add("Authorization", "Bearer "+authToken)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)

		must.Len(chatMessages, numOfMessages)

		for _, m := range chatMessages {
			must.NotEmpty(m.Body)
			must.NotEmpty(m.Id)
			must.NotEmpty(m.Timestamp)
			must.NotEmpty(m.ConvoId)
			must.Equal(userId, m.UserId)
		}
	})

	t.Run("GET /:username/chat-history - has no chat messages", func(t *testing.T) {
		defer clearStorage()

		userId = "rando"

		authToken := login(userId)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-history", userId), nil)
		must.Nil(err)

		req.Header.Add("Authorization", "Bearer "+authToken)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)

		var chatMessages []model.ChatMessage
		err = json.NewDecoder(res.Body).Decode(&chatMessages)
		must.Nil(err)

		must.Empty(chatMessages)
	})

	t.Run("POST /:username/chat-history - successful", func(t *testing.T) {
		defer clearStorage()

		authToken := login(userId)

		body := []byte(`{
			"body": "Hello, world!"
		}`)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-message", userId), bytes.NewBuffer(body))
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+authToken)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusCreated, res.StatusCode)
	})

	t.Run("POST /:username/chat-history - malformed request", func(t *testing.T) {
		defer clearStorage()

		authToken := login(userId)

		body := []byte(`Hello, world!`)
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-message", userId), bytes.NewBuffer(body))
		must.Nil(err)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+authToken)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Invalid/Expired JWT authentication token", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:5000/api/users/%s/chat-history", userId), nil)
		must.Nil(err)

		req.Header.Add("Authorization", "Bearer dummy")

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("POST /register", func(t *testing.T) {
		defer clearStorage()

		body := []byte(`{
			"username": "esia",
			"password": "dummy",
			"email": "esia@mail.com"
		}`)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/register", bytes.NewBuffer(body))
		must.Nil(err)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusCreated, res.StatusCode)

		body = []byte(`{
			"username": "esia",
			"password": "dummy",
			"email": "esia@mail.com"
		}`)
		req, err = http.NewRequest(http.MethodPost, "http://localhost:5000/register", bytes.NewBuffer(body))
		must.Nil(err)

		res, err = http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("POST /login", func(t *testing.T) {
		defer clearStorage()

		body := []byte(`{
			"username": "esia",
			"password": "dummy",
			"email": "esia@mail.com"
		}`)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/register", bytes.NewBuffer(body))
		must.Nil(err)

		res, err := http.DefaultClient.Do(req)
		must.Nil(err)
		must.Equal(http.StatusCreated, res.StatusCode)

		body = []byte(`{
			"username": "esia",
			"password": "dummy"
		}`)
		req, err = http.NewRequest(http.MethodPost, "http://localhost:5000/login", bytes.NewBuffer(body))
		must.Nil(err)

		res, err = http.DefaultClient.Do(req)
		must.Nil(err)
		//must.Equal(http.StatusInternalServerError, res.StatusCode)
	})
}
