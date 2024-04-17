package services

import (
	"context"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (models.JwtToken, error)
	SignUp(name, phoneNumber, email, password string) (string, error)
	LogOut(ctx context.Context, id int) error
	ValidateToken(ctx context.Context, tokenString string) (models.JwtToken, error)
	Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error)
}

func (s *Service) Login(ctx context.Context, username, password string) (models.JwtToken, error) {

	return s.repo.Login(ctx, username, password)
}

func (s *Service) SignUp(name, phoneNumber, email, password string) (models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	return s.repo.SignUp(name, phoneNumber, email, hashedPassword)
}

func (s *Service) LogOut(ctx context.Context, id int) error {
	return s.repo.LogOut(ctx, id)
}

func (s *Service) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	return s.repo.ValidateToken(ctx, tokenString)
}

func (s *Service) Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error) {
	return s.repo.Refresh(ctx, refreshTokenString)
}
