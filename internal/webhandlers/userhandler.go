package webhandlers

import (
	"astroboy/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (wh *WebHandler) GetChatHistory(c echo.Context) error {
	res, err := wh.historyService.ForUser(c.Param("username"))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to get chat history"))
	}

	return c.JSON(http.StatusOK, res)
}

func (wh *WebHandler) AddChatMessage(c echo.Context) error {
	var chatMsg model.ChatMessage
	err := c.Bind(&chatMsg)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to process request"))
	}

	err = wh.historyService.AddChatMessage(c.Param("username"), &chatMsg)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to create chat message"))
	}

	return c.JSON(http.StatusCreated, "successfully created")
}
