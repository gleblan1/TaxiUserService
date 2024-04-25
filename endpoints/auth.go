package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func SignUp(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.RegisterRequest)
		return UserService.SignUp(ctx, requestBody)
	}
}

func Login(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.LoginRequest)
		return UserService.Login(ctx, requestBody)
	}
}

func LogOut(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.LogoutRequest)
		err := UserService.LogOut(ctx, requestBody)
		if err != nil {
			return nil, err
		}
		return UserService.LogOut(ctx, requestBody), nil
	}
}

func Refresh(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.RefreshRequestBody)
		return UserService.Refresh(ctx, requestBody)
	}
}
