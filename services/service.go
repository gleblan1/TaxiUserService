package services

import "github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"

type IService interface {
	IAuthService
	//...Service
	//...
}

type Service struct {
	repo *repositories.Repository
}

func NewServices(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}
