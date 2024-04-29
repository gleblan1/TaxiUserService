package http

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
)

type Handler struct {
	e *endpoints.Endpoints
}

type HandlerOptions func(*Handler)

func NewHandler(options ...HandlerOptions) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithAuthService(e *endpoints.Endpoints) HandlerOptions {
	return func(h *Handler) {
		h.e = e
	}
}
