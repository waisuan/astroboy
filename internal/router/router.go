package router

import (
	custommiddleware "astroboy/internal/router/middlewares"
	"astroboy/internal/webhandlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(wh *webhandlers.WebHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(custommiddleware.Authentication())

	apiGroup := e.Group("/api")

	userGroup := apiGroup.Group("/users")
	userGroup.GET("/:username/chat-history", wh.GetChatHistory)
	userGroup.POST("/:username/chat-message", wh.AddChatMessage)

	return e
}
