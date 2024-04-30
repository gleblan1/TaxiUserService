package repositories

import (
	"github.com/jmoiron/sqlx"
)

type UserPostgresRepository struct {
	db *sqlx.DB
}

type PostgresOption func(*UserPostgresRepository)

func NewPostgres(options ...PostgresOption) *UserPostgresRepository {
	postgres := &UserPostgresRepository{}
	for _, option := range options {
		option(postgres)
	}
	return postgres
}

func WithSqlxDB(db *sqlx.DB) PostgresOption {
	return func(postgres *UserPostgresRepository) {
		postgres.db = db
	}
}
