package webhandlers

import (
	"astroboy/internal/auth"
	"astroboy/internal/chat"
	"astroboy/internal/dependencies"
)

type WebHandler struct {
	Deps           *dependencies.Dependencies
	HistoryService chat.IHistory
	AuthService    auth.IAuth
}

func NewWebHandler(deps *dependencies.Dependencies) *WebHandler {
	return &WebHandler{
		Deps:           deps,
		HistoryService: chat.NewHistoryService(deps),
		AuthService:    auth.NewAuthService(deps),
	}
}
