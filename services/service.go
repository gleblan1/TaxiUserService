package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
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

func NewService(options ...func(*Service)) *Service {
	service := &Service{}
	for _, option := range options {
		option(service)
	}
	return service
}

func WithAuthRepo(repo *repositories.Repository) func(*Service) {
	return func(s *Service) {
		s.repo = repo
	}
}

func (s *Service) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	claims, err := utils.ExtractClaims(tokenString)
	if err != nil {
		return false, err
	}
	userId := claims.Audience
	sessionId := claims.Session
	accessToken, err := s.repo.ValidateToken(ctx, userId, sessionId)
	if err != nil {
		return false, err
	}
	if accessToken == tokenString {
		return true, nil
	}
	return false, nil
}
