package router

import (
	"astroboy/graph"
	"astroboy/internal/router/middlewares"
	"astroboy/internal/webhandlers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(wh *webhandlers.WebHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())

	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL", "/gql/query")

	gqlGroup := e.Group("/gql")
	gqlGroup.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	}, middlewares.Authenticator(wh.Deps.Config.JwtSigningKey))
	gqlGroup.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.POST("/login", wh.Login)
	e.POST("/register", wh.Register)

	apiGroup := e.Group("/api")
	apiGroup.Use(middlewares.Authenticator(wh.Deps.Config.JwtSigningKey))

	userGroup := apiGroup.Group("/users")
	userGroup.GET("/:username/chat-history", wh.GetChatHistory)
	userGroup.POST("/:username/chat-message", wh.AddChatMessage)

	return e
}
