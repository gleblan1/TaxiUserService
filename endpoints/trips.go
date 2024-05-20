package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func RateTrip(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.RateTripRequest)
		return UserService.RateTrip(ctx, requestBody)
	}
}

func GetTripsHistory(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.GetHistoryRequest)
		return UserService.GetTripsHistory(ctx, requestBody)
	}
}
