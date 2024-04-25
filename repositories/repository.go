package repositories

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type UserRepository interface {
	Auth
}

type Auth interface {
	GetData(phone string) (string, int, error)
	SignUp(name, phoneNumber, email, password string) (models.User, error)
	LogOut(ctx context.Context, session, id int) error
	GetRefreshToken(ctx context.Context, id, session string) string
	GetAccessToken(ctx context.Context, id, session string) string
	SetTokens(ctx context.Context, accessToken string, refreshToken, id, session string)
}

type Repository struct {
	db     *sqlx.DB
	client redis.Client
}

func NewRepository(options ...func(*Repository)) *Repository {
	repo := &Repository{}
	for _, option := range options {
		option(repo)
	}
	return repo
}

func WithPostgresRepository(db *sqlx.DB) func(*Repository) {
	return func(r *Repository) {
		r.db = db
	}
}

func WithRedisClient(client redis.Client) func(*Repository) {
	return func(r *Repository) {
		r.client = client
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
