package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

type Endpoints struct {
	SignUp  Endpoint
	Login   Endpoint
	LogOut  Endpoint
	Refresh Endpoint
}

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func MakeEndpoints(UserService services.UserService) *Endpoints {
	return &Endpoints{
		SignUp:  SignUp(UserService),
		Login:   Login(UserService),
		LogOut:  LogOut(UserService),
		Refresh: Refresh(UserService),
	}
}
