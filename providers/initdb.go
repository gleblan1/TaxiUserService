package providers

import (
	"context"
	"fmt"

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
