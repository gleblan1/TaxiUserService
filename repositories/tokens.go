package repositories

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
)

func (r *Repository) GetRefreshToken(ctx context.Context, id, session string) string {
	return strings.Split(r.client.Client.Get(ctx, id+"."+session).String(), " ")[3]
}

func (r *Repository) GetAccessToken(ctx context.Context, id, session string) string {
	return strings.Split(r.client.Client.Get(ctx, id+"."+session).String(), " ")[2]
}

func (r *Repository) SetTokens(ctx context.Context, accessToken string, refreshToken, id, session string) error {
	err := r.client.Client.Set(ctx, id+"."+session, accessToken+" "+refreshToken, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ValidateToken(ctx context.Context, userId, sessionId string) (string, error) {
	intId, _ := strconv.Atoi(userId)
	if r.CheckIsUserDeleted(intId) {
		return "", errors.New("user has been deleted")
	}
	tokenFromRedis := strings.Split(r.client.Client.Get(ctx, userId+"."+sessionId).String(), " ")[2]
	return tokenFromRedis, nil
}
