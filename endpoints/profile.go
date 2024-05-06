package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func GetAccountInfo(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.GetAccountInfoRequest)
		return UserService.GetAccountInfo(ctx, requestBody)
	}
}

func UpdateProfile(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.UpdateProfileRequest)
		return UserService.UpdateProfile(ctx, requestBody)
	}
}

func DeleteProfile(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.DeleteProfileRequest)
		return UserService.DeleteProfile(ctx, requestBody), nil
	}
}
