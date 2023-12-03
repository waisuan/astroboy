package webhandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (wh *WebHandler) GetChatHistory(c echo.Context) error {
	res, err := wh.historyService.ForUser(c.Param("username"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CustomErrorResponse("failed to process request"))
	}

	return c.JSON(http.StatusOK, res)
}
