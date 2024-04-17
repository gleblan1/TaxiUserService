package services

import "github.com/GO-Trainee/GlebL-innotaxi-userservice/models"

type IProfileService interface {
	GetAccountInfo(id int) models.User
	UpdateProfile(id int, newData models.User) models.User
	DeleteProfile(id int)
}

func (s *Service) GetAccountInfo(id int) models.User {
	var user models.User
	return user
}

func (s *Service) UpdateProfile(id int, newData models.User) models.User {
	var user models.User
	return user
}

func (s *Service) DeleteProfile(id int) {
}
