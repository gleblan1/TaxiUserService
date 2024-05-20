package services

import (
	"context"
	"fmt"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
)

func (s *Service) RateTrip(ctx context.Context, req requests.RateTripRequest) (models.Trip, error) {
	user, err := s.repo.GetUserById(ctx, req.UserId)
	lastTrip, err := s.repo.GetLastTrip(ctx, user)
	if err != nil {
		return models.Trip{}, err
	}
	_, err = s.repo.RateTrip(ctx, lastTrip.Id, req.Rate)
	if err != nil {
		return models.Trip{}, err
	}
	result, err := s.repo.GetLastTrip(ctx, user)
	return result, nil
}

func (s *Service) GetTripsHistory(ctx context.Context, req requests.GetHistoryRequest) (models.TripHistory, error) {
	user, err := s.repo.GetUserById(ctx, req.UserId)
	fmt.Println(user)
	if err != nil {
		return models.TripHistory{}, err
	}
	history, err := s.repo.GetTripsHistory(ctx, user)
	if err != nil {
		return models.TripHistory{}, err
	}

	response := models.TripHistory{
		Trips: history,
		User:  user,
	}

	return response, nil
}
