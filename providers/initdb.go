package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	Psql  *sqlx.DB
	Redis *redis.Client
}

func NewDB(options ...func(db *DB)) *DB {
	db := &DB{}
	for _, option := range options {
		option(db)
	}
	return db
}

func WithPsql(psql *sqlx.DB) func(db *DB) {
	return func(db *DB) {
		db.Psql = psql
	}
}

func WithRedis(redis *redis.Client) func(db *DB) {
	return func(db *DB) {
		db.Redis = redis
	}
}
