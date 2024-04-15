package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type IRepository interface {
	IAuthRepository
}

type Repository struct {
	AuthRepository
}

func NewRepository(db *sqlx.DB, client redis.Client) *Repository {
	return &Repository{AuthRepository: AuthRepository{
		db,
		client,
	}}
}
