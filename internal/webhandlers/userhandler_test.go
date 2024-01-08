package webhandlers

import (
	"astroboy/internal/mocks"
	"astroboy/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHandler_GetChatHistory(t *testing.T) {
	t.Run("runs", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:username/chat-history")
		c.SetParamNames("username")
		c.SetParamValues("esia")

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockChatHistory := mocks.NewMockIHistory(mockCtrl)
		mockChatHistory.EXPECT().ForUser("esia").Return([]model.ChatMessage{}, nil)

		h := &WebHandler{
			HistoryService: mockChatHistory,
		}

		if assert.NoError(t, h.GetChatHistory(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
