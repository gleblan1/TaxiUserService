package repositories

import "github.com/GO-Trainee/GlebL-innotaxi-userservice/models"

type IProfileRepository interface {
	GetAccountInfo(id int) models.User
	UpdateProfile(id int, newData models.User) models.User
	DeleteProfile(id int)
}
