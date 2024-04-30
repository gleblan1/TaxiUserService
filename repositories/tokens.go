package repositories

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	userIsDeletedErr = errors.New("user is deleted")
	tokenNotFoundErr = errors.New("token not found")
)

func (r *Repository) GetRefreshToken(ctx context.Context, id, session string) string {
	return strings.Split(r.redis.Get(ctx, id+"."+session).String(), " ")[3]
}

func (r *Repository) FindTokens(ctx context.Context, redisKey string) (int64, error) {
	return r.redis.Exists(ctx, redisKey).Result()
}

func (r *Repository) DeleteTokens(ctx context.Context, redisKey string) {
	r.redis.Del(ctx, redisKey)
}

func (r *Repository) GetAccessToken(ctx context.Context, id, session string) string {
	return strings.Split(r.redis.Get(ctx, id+"."+session).String(), " ")[2]
}

func (r *Repository) SetTokens(ctx context.Context, accessToken string, refreshToken, id, session string) error {
	err := r.redis.Set(ctx, id+"."+session, accessToken+" "+refreshToken, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ValidateToken(ctx context.Context, userId, sessionId string) (string, error) {
	intId, _ := strconv.Atoi(userId)
	if r.IsUserDeleted(ctx, intId) {
		return "", userIsDeletedErr
	}
	tokens := strings.Split(r.redis.Get(ctx, userId+"."+sessionId).String(), " ")
	if len(tokens) == 0 {
		return "", tokenNotFoundErr
	}
	tokenFromRedis := r.GetAccessToken(ctx, userId, sessionId)
	return tokenFromRedis, nil
}
