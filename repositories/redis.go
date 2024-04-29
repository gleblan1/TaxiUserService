package repositories

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

type RedisOption func(*Redis)

func NewRedis(options ...RedisOption) *Redis {
	redisStore := &Redis{}
	for _, option := range options {
		option(redisStore)
	}
	return redisStore
}

func WithRedisClient(client *redis.Client) RedisOption {
	return func(redis *Redis) {
		redis.Client = client
	}
}
