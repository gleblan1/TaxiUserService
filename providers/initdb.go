package providers

import (
	"context"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func InitDB() (postgresDb *sqlx.DB, redisDb *redis.Client, err error) {
	postgresDb, err = initPostgres()
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres db err: %s", err.Error())
	}
	redisDb, err = initRedis(context.TODO())
	if err != nil {
		return nil, nil, fmt.Errorf("init redis db err: %s", err.Error())
	}
	return postgresDb, redisDb, nil
}

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

func initRedis(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error: %w", err)
	}

	return client, nil
}
