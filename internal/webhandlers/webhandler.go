package webhandlers

import (
	"astroboy/internal/chat"
	"astroboy/internal/dependencies"
)

type WebHandler struct {
	deps           *dependencies.Dependencies
	historyService *chat.HistoryService
}

func NewWebHandler(deps *dependencies.Dependencies) *WebHandler {
	return &WebHandler{
		deps:           deps,
		historyService: chat.NewHistoryService(deps),
	}
}
