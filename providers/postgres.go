package providers

import (
	"fmt"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/jmoiron/sqlx"
)

func initPostgres() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", utils.DbConnectionString())
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}
	return db, nil
}
