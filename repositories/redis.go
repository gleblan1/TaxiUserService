package repositories

import (
	"github.com/redis/go-redis/v9"
)

type UserRedisRepository struct {
	Client *redis.Client
}

type RedisOption func(*UserRedisRepository)

func NewRedis(options ...RedisOption) *UserRedisRepository {
	redisStore := &UserRedisRepository{}
	for _, option := range options {
		option(redisStore)
	}
	return redisStore
}

func WithRedisClient(client *redis.Client) RedisOption {
	return func(redis *UserRedisRepository) {
		redis.Client = client
	}
}
