package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type UserFromDb struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Password    string `db:"password_hash"`
	Rating      string `db:"rating"`
}

func (r *Repository) GetData(phone string) (string, int, error) {
	var userId int
	err := r.db.db.QueryRowx("SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phone).Scan(&userId)
	if err != nil {
		return "", userId, errors.New("user not found")
	}

	var passwordFromDB string
	err = r.db.db.QueryRow("SELECT password_hash FROM users WHERE id = $1 AND deleted_at IS NULL", userId).Scan(&passwordFromDB)
	if err != nil {
		return "", userId, err
	}
	return passwordFromDB, userId, nil
}

func (r *Repository) CreateUser(name, phoneNumber, email, password string) (int, error) {
	var userId int
	stmt, err := r.db.db.Prepare("INSERT INTO users(name, phone_number, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id")
	if err != nil {
		fmt.Println(1)
		return 0, fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
		}
	}(stmt)

	err = stmt.QueryRow(name, phoneNumber, email, password).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %w", err)
	}
	return userId, nil
}

func (r *Repository) GetUser(id int) (models.User, error) {
	var userFromDb UserFromDb
	err := r.db.db.QueryRow("SELECT id, name, phone_number, email, password_hash, rating FROM users WHERE id=$1", id).Scan(&userFromDb.Id, &userFromDb.Name, &userFromDb.PhoneNumber, &userFromDb.Email, &userFromDb.Password, &userFromDb.Rating)
	if err != nil {
		fmt.Println(2, err)
		return models.User{}, err
	}
	rating, err := strconv.ParseFloat(userFromDb.Rating, 32)

	return models.User{
		Id:          userFromDb.Id,
		Name:        userFromDb.Name,
		PhoneNumber: userFromDb.PhoneNumber,
		Email:       userFromDb.Email,
		Password:    userFromDb.Password,
		Rating:      float32(rating),
	}, nil
}

func (r *Repository) LogOut(ctx context.Context, session, id int) error {
	exists, err := r.client.Client.Exists(ctx, strconv.Itoa(id)+"."+strconv.Itoa(session)).Result()
	if err != nil {
		return fmt.Errorf("log out: %w", err)
	}
	if exists == 1 {
		r.client.Client.Del(context.Background(), strconv.Itoa(id)+"."+strconv.Itoa(session))
	} else {
		return fmt.Errorf("log out: %w", errors.New("already log out"))
	}
	return nil
}

func (r *Repository) CheckUserData(name, phoneNumber, email string) error {
	existingUser := models.User{}

	err := r.db.db.Get(&existingUser, "SELECT name FROM users WHERE name = $1 AND deleted_at IS NULL", name)
	if err == nil {
		return errors.New("name already in use")
	}
	err = r.db.db.Get(&existingUser.Name, "SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
	if err == nil {
		return errors.New("phone already in use")
	}
	err = r.db.db.Get(&existingUser, "SELECT id FROM users WHERE email = $1 AND deleted_at IS NULL", email)
	if err == nil {
		return errors.New("email already in use")
	}

	return nil
}

func (r *Repository) CheckIsUserDeleted(userId int) bool {
	var deletedAt time.Time
	err := r.db.db.Get(&deletedAt, "SELECT deleted_at FROM users WHERE id = $1 AND deleted_at IS NULL ", userId)
	if err != nil {
		return false
	}
	return deletedAt != time.Unix(0, 0)
}
