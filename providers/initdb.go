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
	redisDb = initRedis(context.TODO())
	return postgresDb, redisDb, nil
}

func initPostgres() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", utils.DbConnectionString())
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping error: %w", err)
	}
	//
	return db, nil
}

func initRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	return client
}
