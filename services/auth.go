package services

import (
	"context"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (models.JwtToken, error)
	SignUp(name, phoneNumber, email, password string) (string, error)
	LogOut(id int) error
	ValidateToken(tokenString string) (models.JwtToken, error)
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

func (s *Service) LogOut(id int) error {
	return s.repo.LogOut(id)
}
