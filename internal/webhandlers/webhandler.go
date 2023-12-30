package webhandlers

import (
	"astroboy/internal/auth"
	"astroboy/internal/chat"
	"astroboy/internal/dependencies"
)

type WebHandler struct {
	deps           *dependencies.Dependencies
	historyService *chat.HistoryService
	authService    *auth.AuthService
}

func NewWebHandler(deps *dependencies.Dependencies) *WebHandler {
	return &WebHandler{
		deps:           deps,
		historyService: chat.NewHistoryService(deps),
		authService:    auth.NewAuthService(deps),
	}
}
