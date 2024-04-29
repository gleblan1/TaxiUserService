package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

type PostgresOption func(*Postgres)

func NewPostgres(options ...PostgresOption) *Postgres {
	postgres := &Postgres{}
	for _, option := range options {
		option(postgres)
	}
	return postgres
}

func WithSqlxDB(db *sqlx.DB) PostgresOption {
	return func(postgres *Postgres) {
		postgres.db = db
	}
}
