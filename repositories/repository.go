package repositories

import (
	"context"
	"errors"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type IRepository interface {
	IAuthRepository
}

type Repository struct {
	db     *sqlx.DB
	client redis.Client
}

func NewRepository(db *sqlx.DB, client redis.Client) *Repository {
	return &Repository{
		db,
		client,
	}
}

func (r *Repository) CheckUserData(name, phoneNumber, email string) error {
	existingUser := models.User{}

	err := r.db.Get(&existingUser, "SELECT name FROM users WHERE name = $1 AND deleted_at IS NULL", name)
	if err == nil {
		return errors.New("name already in use")
	}
	err = r.db.Get(&existingUser.Name, "SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
	if err == nil {
		return errors.New("phone already in use")
	}
	err = r.db.Get(&existingUser, "SELECT id FROM users WHERE email = $1 AND deleted_at IS NULL", email)
	if err == nil {
		return errors.New("email already in use")
	}

	return nil
}

func (r *Repository) CheckIsUserDeleted(userId int) bool {
	var deletedAt time.Time
	err := r.db.Get(&deletedAt, "SELECT deleted_at FROM users WHERE id = $1 AND deleted_at IS NULL ", userId)
	if err != nil {
		return false
	}
	return deletedAt != time.Unix(0, 0)
}

func (r *Repository) ValidateToken(ctx context.Context, userId, sessionId string) (string, error) {
	intId, _ := strconv.Atoi(userId)
	if r.CheckIsUserDeleted(intId) {
		return "", errors.New("user has been deleted")
	}
	tokenFromRedis := strings.Split(r.client.Get(ctx, userId+"."+sessionId).String(), " ")[2]
	return tokenFromRedis, nil
}
