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
	"time"
)

type IAuthRepository interface {
	Login(ctx context.Context, username, password string) (models.JwtToken, error)
	SignUp(username, email, password string) (string, error)
	LogOut(id int) error
	ValidateToken(username, password string) (bool, error)
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
		return models.JwtToken{}, errors.New("Wrong username")
	}

	var passwordFromDB string
	err = r.db.QueryRow("SELECT password_hash FROM users WHERE id = $1", userId).Scan(&passwordFromDB)
	if err != nil {
		return models.JwtToken{}, err
	}
	isPasswordCorrect, err := utils.ComparePassword(passwordFromDB, password)
	if err != nil {
		return models.JwtToken{}, errors.New("Wrong password")
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

func (r *Repository) LogOut(id int) error {
	var userId int

	userId = id
	exists, err := r.client.Exists(context.TODO(), strconv.Itoa(userId)).Result()
	if err != nil {
		return fmt.Errorf("log out: %w", err)
	}
	if exists == 1 {
		r.client.Del(context.Background(), strconv.Itoa(userId))
	} else {
		return fmt.Errorf("log out: %w", errors.New("already log out"))
	}
//clean up the code 
//check all the error handlers (bug: i have a few errors in response)
//implement validate token
//fix bug with log out (bug: user can log out using any token)
	return nil
}

func (r *Repository) ValidateToken(username, password string) (bool, error) {
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
