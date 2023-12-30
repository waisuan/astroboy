package webhandlers

import (
	"astroboy/internal/auth"
	"astroboy/internal/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (wh *WebHandler) GetChatHistory(c echo.Context) error {
	res, err := wh.historyService.ForUser(c.Param("username"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch chat history")
	}

	return c.JSON(http.StatusOK, res)
}

func (wh *WebHandler) AddChatMessage(c echo.Context) error {
	var chatMsg model.ChatMessage
	err := c.Bind(&chatMsg)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	err = wh.historyService.AddChatMessage(c.Param("username"), &chatMsg)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to add chat message")
	}

	return c.JSON(http.StatusCreated, "successfully created")
}

func (wh *WebHandler) Login(c echo.Context) error {
	creds := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	username := creds["username"]
	//password := creds["password"]
	// TODO: properly authenticate against existing user

	token, err := auth.GenerateJwtToken(username.(string), wh.deps.Config.JwtSigningKey)
	if err != nil {
		log.Error(err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (wh *WebHandler) Register(c echo.Context) error {
	creds := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	username := creds["username"].(string)
	password := creds["password"].(string)
	email := creds["email"].(string)

	err = wh.authService.RegisterUser(username, password, email)
	if err != nil {
		log.Error(err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, nil)
}
