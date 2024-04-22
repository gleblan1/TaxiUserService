package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type UserService interface {
	Auth
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
