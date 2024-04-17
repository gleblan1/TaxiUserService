package handler

import (
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

type Handler struct {
	s *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{s: s}
}
