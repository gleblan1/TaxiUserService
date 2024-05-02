package repositories

import (
	"context"
	"fmt"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type TripModel struct {
	id       int    `db:"id"`
	taxiType int    `db:"taxi_type"`
	from     string `db:"from_address"`
	to       string `db:"to_address"`
	rate     int    `db:"rate"`
}

func (r *Repository) GetLastTrip(ctx context.Context, user models.User) (models.Trip, error) {
	trip := TripModel{}
	err := r.db.QueryRow("SELECT id, taxi_type, from_address, to_address, rate FROM trips WHERE user_id=$1 ORDER BY id DESC LIMIT 1", 32).Scan(&trip.id, &trip.taxiType, &trip.from, &trip.to, &trip.rate)
	if err != nil {
		fmt.Println(trip)
		return models.Trip{}, err
	}

	result := models.Trip{
		Id:       trip.id,
		TaxiType: trip.taxiType,
		From:     trip.from,
		To:       trip.to,
		Rate:     trip.rate,
		User:     user,
	}
	fmt.Println(trip)
	return result, nil
}

func (r *Repository) RateTrip(ctx context.Context, tripId int, rate int) (float32, error) {
	_, err := r.db.Exec("UPDATE trips SET rate=$1 WHERE id=$2", rate, tripId)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *Repository) GetTripsHistory(ctx context.Context, user models.User) ([]models.Trip, error) {
	var result []models.Trip
	rows, err := r.db.Query("SELECT id, taxi_type, from_address, to_address, rate FROM trips WHERE user_id=$1 ORDER BY id DESC LIMIT 1", 32)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		trip := TripModel{}
		resultItem := models.Trip{}
		rows.Scan(&trip.id, &trip.taxiType, &trip.from, &trip.to, &trip.rate)
		resultItem = models.Trip{
			Id:       trip.id,
			TaxiType: trip.taxiType,
			From:     trip.from,
			To:       trip.to,
			Rate:     trip.rate,
		}
		result = append(result, resultItem)
	}
	return result, nil
}
