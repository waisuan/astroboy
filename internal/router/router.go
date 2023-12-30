package router

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/router/middlewares"
	"astroboy/internal/webhandlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(wh *webhandlers.WebHandler, cfg *dependencies.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())

	e.POST("/login", wh.Login)
	e.POST("/register", wh.Register)

	apiGroup := e.Group("/api")
	apiGroup.Use(middlewares.Authenticator(cfg.JwtSigningKey))

	userGroup := apiGroup.Group("/users")
	userGroup.GET("/:username/chat-history", wh.GetChatHistory)
	userGroup.POST("/:username/chat-message", wh.AddChatMessage)

	return e
}
