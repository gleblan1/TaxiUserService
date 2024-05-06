package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
)

type UserService interface {
	Auth
	Profile
}

type Auth interface {
	SignIn(ctx context.Context, requestBody requests.SignInRequest) (models.JwtToken, error)
	SignUp(ctx context.Context, requestBody requests.RegisterRequest) (models.User, error)
	SignOut(ctx context.Context, request requests.LogoutRequest) error
	ValidateToken(ctx context.Context, tokenString string) (bool, error)
	RefreshTokens(ctx context.Context, requestBody requests.RefreshTokensRequest) (models.JwtToken, error)
}

type Profile interface {
	GetAccountInfo(ctx context.Context, req requests.GetAccountInfoRequest) (models.User, error)
	UpdateProfile(ctx context.Context, req requests.UpdateProfileRequest) (models.User, error)
	DeleteProfile(ctx context.Context, req requests.DeleteProfileRequest) error
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

func WithRepo(repo *repositories.Repository) ServiceOption {
	return func(s *Service) {
		s.repo = repo
	}
}
