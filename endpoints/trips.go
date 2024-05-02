package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func RateTrip(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.RateTripRequest)
		return UserService.RateTrip(ctx, requestBody)
	}
}

func GetTripsHistory(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.GetHistoryRequest)
		return UserService.GetTripsHistory(ctx, requestBody)
	}
}
