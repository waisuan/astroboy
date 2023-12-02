package webhandlers

import (
	"astroboy/internal/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (wh *WebHandler) GetUser(c echo.Context) error {
	log.Printf("GetUser: %s\n", c.Param("username"))

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

func (wh *WebHandler) GetChatHistory(c echo.Context) error {
	res, err := wh.historyService.ForUser(c.Param("username"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to process request"))
	}

	return c.JSON(http.StatusOK, res)
}
