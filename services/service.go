package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
)

type UserService interface {
	Auth
}

type Auth interface {
	Login(ctx context.Context, requestBody requests.LoginRequest) (models.JwtToken, error)
	SignUp(ctx context.Context, requestBody requests.RegisterRequest) (models.User, error)
	LogOut(ctx context.Context, request requests.LogoutRequest) error
	ValidateToken(ctx context.Context, tokenString string) (bool, error)
	Refresh(ctx context.Context, requestBody requests.RefreshRequestBody) (models.JwtToken, error)
}

type Service struct {
	repo *repositories.Repository
}

type ServiceOption func(service *Service)

func NewService(options ...ServiceOption) *Service {
	service := &Service{}
	for _, option := range options {
		option(service)
	}
	return service
}

func WithAuthRepo(repo *repositories.Repository) ServiceOption {
	return func(s *Service) {
		s.repo = repo
	}
}
