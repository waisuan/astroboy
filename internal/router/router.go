package router

import (
	"astroboy/internal/webhandlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(wh *webhandlers.WebHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	// TODO: Auth

	apiGroup := e.Group("/api")

	userGroup := apiGroup.Group("/users")
	userGroup.GET("/:username", wh.GetUser)
	userGroup.GET("/:username/chat-history", wh.GetChatHistory)

	return e
}
