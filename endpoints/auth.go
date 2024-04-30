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

func SignIn(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.SignInRequest)
		return UserService.SignIn(ctx, requestBody)
	}
}

func SignOut(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.LogoutRequest)
		err := UserService.SignOut(ctx, requestBody)
		if err != nil {
			return nil, err
		}
		return UserService.SignOut(ctx, requestBody), nil
	}
}

func RefreshTokens(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.RefreshTokensRequest)
		return UserService.RefreshTokens(ctx, requestBody)
	}
}
