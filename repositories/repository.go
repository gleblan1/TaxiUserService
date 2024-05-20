package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	db    *sqlx.DB
	redis redis.Client
}

type ReposOption func(*Repository)

func NewRepository(options ...ReposOption) *Repository {
	repo := &Repository{}
	for _, option := range options {
		option(repo)
	}
	return repo
}

func WithPostgresRepository(db UserPostgresRepository) ReposOption {
	return func(r *Repository) {
		r.db = db.db
	}
}

func WithRedis(client UserRedisRepository) ReposOption {
	return func(r *Repository) {
		r.redis = *client.Client
	}
}
