package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type Auth interface {
	GetData(phone string) (string, int, error)
	SignUp(name, phoneNumber, email, password string) (models.User, error)
	LogOut(ctx context.Context, session, id int) error
	GetRefreshToken(ctx context.Context, id, session string) string
	GetAccessToken(ctx context.Context, id, session string) string
	SetTokens(ctx context.Context, accessToken string, refreshToken, id, session string)
}

func (r *Repository) GetData(phone string) (string, int, error) {
	var userId int
	err := r.db.QueryRowx("SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phone).Scan(&userId)
	if err != nil {
		return "", userId, errors.New("user not found")
	}

	var passwordFromDB string
	err = r.db.QueryRow("SELECT password_hash FROM users WHERE id = $1 AND deleted_at IS NULL", userId).Scan(&passwordFromDB)
	if err != nil {
		return "", userId, err
	}
	return passwordFromDB, userId, nil
}

func (r *Repository) SignUp(name, phoneNumber, email, password string) (models.User, error) {
	user := models.User{
		Name:        name,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    password,
	}
	stmt, err := r.db.Prepare("INSERT INTO users(name, phone_number, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id")
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

func (r *Repository) LogOut(ctx context.Context, session, id int) error {
	exists, err := r.client.Exists(ctx, strconv.Itoa(id)+"."+strconv.Itoa(session)).Result()
	if err != nil {
		return fmt.Errorf("log out: %w", err)
	}
	if exists == 1 {
		fmt.Println(id, session)
		r.client.Del(context.Background(), strconv.Itoa(id)+"."+strconv.Itoa(session))
	} else {
		return fmt.Errorf("log out: %w", errors.New("already log out"))
	}
	return nil
}

func (r *Repository) GetRefreshToken(ctx context.Context, id, session string) string {
	return strings.Split(r.client.Get(ctx, id+"."+session).String(), " ")[3]
}

func (r *Repository) GetAccessToken(ctx context.Context, id, session string) string {
	return strings.Split(r.client.Get(ctx, id+"."+session).String(), " ")[2]
}

func (r *Repository) SetTokens(ctx context.Context, accessToken string, refreshToken, id, session string) error {
	err := r.client.Set(ctx, id+"."+session, accessToken+" "+refreshToken, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
