package webhandler

import "astroboy/internal/dependencies"

type WebHandler struct {
	deps *dependencies.Dependencies
}

func NewWebHandler(deps *dependencies.Dependencies) *WebHandler {
	return &WebHandler{
		deps: deps,
	}
}
