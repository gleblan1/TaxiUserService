package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

type Endpoints struct {
	SignUp        Endpoint
	SignIn        Endpoint
	SignOut       Endpoint
	RefreshTokens Endpoint
}

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func MakeEndpoints(UserService services.UserService) *Endpoints {
	return &Endpoints{
		SignUp:        SignUp(UserService),
		SignIn:        SignIn(UserService),
		SignOut:       SignOut(UserService),
		RefreshTokens: RefreshTokens(UserService),
	}
}
