package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func GetAccountInfo(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.GetAccountInfoRequest)
		return UserService.GetAccountInfo(ctx, requestBody)
	}
}

func UpdateProfile(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.UpdateProfileRequest)
		return UserService.UpdateProfile(ctx, requestBody)
	}
}

func DeleteProfile(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.DeleteProfileRequest)
		return UserService.DeleteProfile(ctx, requestBody), nil
	}
}
