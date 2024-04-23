package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func SignUp(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.RegisterRequest)
		return UserService.SignUp(ctx, requestBody)
	}
}

func Login(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.LoginRequest)
		return UserService.Login(ctx, requestBody)
	}
}

func LogOut(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.LogoutRequest)
		err := UserService.LogOut(ctx, requestBody)
		if err != nil {
			return nil, err
		}
		return UserService.LogOut(ctx, requestBody), nil
	}
}

func Refresh(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.RefreshRequestBody)
		return UserService.Refresh(ctx, requestBody)
	}
}
