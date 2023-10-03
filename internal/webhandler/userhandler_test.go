package webhandler

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	testify "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHandler_GetUser(t *testing.T) {
	assert := testify.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:username")
	c.SetParamNames("username")
	c.SetParamValues("esia")

	t.Run("user exists", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		fakeUser := model.User{
			Username:    "esia",
			Email:       "e-sia@outlook.com",
			DateOfBirth: "12/11/1991",
		}
		j, _ := json.Marshal(fakeUser)
		mockCache := dependencies.NewMockICache(mockCtrl)
		mockCache.EXPECT().Get("esia").Return(string(j), nil)

		wh := NewWebHandler(&dependencies.Dependencies{
			CacheCli: mockCache,
		})

		assert.NoError(wh.GetUser(c))
		assert.Equal(http.StatusOK, rec.Code)

		var resUser model.User
		_ = json.Unmarshal([]byte(rec.Body.String()), &resUser)
		assert.Equal(fakeUser.Username, resUser.Username)
		assert.Equal(fakeUser.Email, resUser.Email)
		assert.Equal(fakeUser.DateOfBirth, resUser.DateOfBirth)
	})

	t.Run("user does not exist", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCache := dependencies.NewMockICache(mockCtrl)
		mockCache.EXPECT().Get("esia").Return("", nil)

		wh := NewWebHandler(&dependencies.Dependencies{
			CacheCli: mockCache,
		})

		assert.NoError(wh.GetUser(c))
		assert.Equal(http.StatusNotFound, rec.Code)
	})

	t.Run("unable to parse the response", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockCache := dependencies.NewMockICache(mockCtrl)
		mockCache.EXPECT().Get("esia").Return("something_went_wrong_here", nil)

		wh := NewWebHandler(&dependencies.Dependencies{
			CacheCli: mockCache,
		})

		assert.NoError(wh.GetUser(c))
		assert.Equal(http.StatusInternalServerError, rec.Code)
	})
}
