package webhandlers

import (
	"astroboy/internal/mocks"
	"astroboy/internal/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHandler_GetChatHistory(t *testing.T) {
	t.Run("returns a 200 response if there are no errors", func(t *testing.T) {
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

		err := h.GetChatHistory(c)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("returns a 500 response if there are errors", func(t *testing.T) {
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
		mockChatHistory.EXPECT().ForUser("esia").Return(nil, errors.New("something bad happened"))

		h := &WebHandler{
			HistoryService: mockChatHistory,
		}

		err := h.GetChatHistory(c)
		assert.NotNil(t, err)

		res, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
