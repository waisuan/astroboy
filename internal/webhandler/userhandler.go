package webhandler

import (
	"astroboy/internal/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (wh *WebHandler) GetUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, CustomErrorResponse("bad request"))
	}

	res, err := wh.deps.CacheCli.Get(user.Username)
	if res == "" || err != nil {
		return c.JSON(http.StatusNotFound, CustomErrorResponse("user does not exist"))
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to process request"))
	}

	return c.JSON(http.StatusOK, user)
}
