package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type IAuthRepository interface {
	Login(ctx context.Context, username, password string) (models.JwtToken, error)
	SignUp(username, email, password string) (string, error)
	LogOut(ctx context.Context, id int) error
	ValidateToken(ctx context.Context, tokenString string) (bool, error)
	Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error)
}

type AuthRepository struct {
	db     *sqlx.DB
	client redis.Client
}

func (r *Repository) Login(ctx context.Context, username, password string) (models.JwtToken, error) {
	jwtToken := models.JwtToken{}

	var userId int

	err := r.db.QueryRowx("SELECT id FROM users WHERE name = $1", username).Scan(&userId)
	if err != nil {
		return models.JwtToken{}, errors.New("user not found")
	}

	var passwordFromDB string
	err = r.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", userId).Scan(&passwordFromDB)
	if err != nil {
		return models.JwtToken{}, err
	}
	isPasswordCorrect, err := utils.ComparePassword(passwordFromDB, password)
	if err != nil {
		return models.JwtToken{}, errors.New("wrong password")
	}

	if isPasswordCorrect {
		accessToken, refreshToken, err := utils.GenerateTokens(strconv.Itoa(userId))
		if err != nil {
			return models.JwtToken{}, err
		}
		jwtToken.AccessToken = accessToken
		jwtToken.RefreshToken = refreshToken
		err = r.client.Set(ctx, strconv.Itoa(userId), accessToken+" "+refreshToken, 24*time.Hour).Err()
		if err != nil {
			return models.JwtToken{}, err
		}
		return jwtToken, nil
	}
	return models.JwtToken{}, nil
}

func (r *Repository) SignUp(name, phoneNumber, email, password string) (models.User, error) {
	user := models.User{
		Name:        name,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    password,
	}
	if r.db == nil {
		return user, errors.New("database connection is nil")
	}

	existingUserErr := r.checkUserData(user.Name, user.PhoneNumber, user.Email)
	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot create user: %w", existingUserErr)
	}
	stmt, err := r.db.Prepare("INSERT INTO users(name, phone_number, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return user, fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
		}
	}(stmt)

	err = stmt.QueryRow(name, phoneNumber, email, password).Scan(&user.Id)
	if err != nil {
		return user, fmt.Errorf("error executing SQL statement: %w", err)
	}

	return user, nil
}

func (r *Repository) LogOut(ctx context.Context, id int) error {
	var userId int

	userId = id
	exists, err := r.client.Exists(ctx, strconv.Itoa(userId)).Result()
	if err != nil {
		return fmt.Errorf("log out: %w", err)
	}
	if exists == 1 {
		r.client.Del(context.Background(), strconv.Itoa(userId))
	} else {
		return fmt.Errorf("log out: %w", errors.New("already log out"))
	}
	return nil
}

func (r *Repository) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	claims, err := utils.ExtractClaims(tokenString)
	if err != nil {
		return false, err
	}
	userId := claims.Audience
	tokenFromRedis := strings.Split(r.client.Get(ctx, userId).String(), " ")[2]
	if tokenFromRedis == tokenString {
		return true, nil
	}
	return false, nil
}

func (r *Repository) checkUserData(name, phoneNumber, email string) error {
	existingUser := models.User{}
	err := r.db.Get(&existingUser, "SELECT name FROM users WHERE name = $1", name)
	if err == nil {
		return errors.New("name already in use")
	}
	err = r.db.Get(&existingUser, "SELECT id FROM users WHERE phone_number = $1", phoneNumber)
	if err == nil {
		return errors.New("phone already in use")
	}
	err = r.db.Get(&existingUser, "SELECT id FROM users WHERE email = $1", email)
	if err == nil {
		return errors.New("email already in use")
	}
	return nil
}

func (r *Repository) Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error) {
	token := models.JwtToken{}
	claims, err := utils.ExtractClaims(refreshTokenString)
	if err != nil {
		return token, err
	}
	userId := claims.Audience
	tokensFromRedis := strings.Split(r.client.Get(ctx, userId).String(), " ")[3]
	if tokensFromRedis == refreshTokenString {

		accessToken, refreshToken, err := utils.GenerateTokens(userId)

		if err != nil {
			return models.JwtToken{}, err
		}

		token.AccessToken = accessToken
		token.RefreshToken = refreshToken

		err = r.client.Set(ctx, userId, accessToken+" "+refreshToken, 24*time.Hour).Err()
		if err != nil {
			return models.JwtToken{}, err
		}
	} else {
		return models.JwtToken{}, errors.New("refresh token is invalid")
	}
	return token, nil
}
