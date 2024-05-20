package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

var (
	nameAlreadyInUseErr  = errors.New("username is already in use")
	phoneAlreadyInUseErr = errors.New("phone is already in use")
	emailAlreadyInUseErr = errors.New("email is already in use")
)

type userModel struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	Email       string  `db:"email"`
	PhoneNumber string  `db:"phone_number"`
	Password    string  `db:"password_hash"`
	Rating      float32 `db:"rating"`
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (models.User, error) {
	var userData userModel
	err := r.db.QueryRowx("SELECT id, name, email, phone_number, password_hash, rating FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phone).Scan(&userData.Id, &userData.Name, &userData.Email, &userData.PhoneNumber, &userData.Password, &userData.Rating)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Id:          userData.Id,
		Name:        userData.Name,
		Email:       userData.Email,
		PhoneNumber: userData.PhoneNumber,
		Password:    userData.Password,
		Rating:      userData.Rating,
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (int, error) {
	var userId int
	stmt, err := r.db.Prepare("INSERT INTO users(name, phone_number, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
		}
	}(stmt)

	err = stmt.QueryRow(user.Name, user.PhoneNumber, user.Email, user.Password).Scan(&userId)
	if err != nil {
		return 0, fmt.Errorf("error executing SQL statement: %w", err)
	}
	return userId, nil
}

func (r *Repository) GetUserById(ctx context.Context, id int) (models.User, error) {
	var userFromDb userModel
	err := r.db.QueryRow("SELECT id, name, phone_number, email, password_hash, rating FROM users WHERE id=$1", id).Scan(&userFromDb.Id, &userFromDb.Name, &userFromDb.PhoneNumber, &userFromDb.Email, &userFromDb.Password, &userFromDb.Rating)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:          userFromDb.Id,
		Name:        userFromDb.Name,
		PhoneNumber: userFromDb.PhoneNumber,
		Email:       userFromDb.Email,
		Password:    userFromDb.Password,
		Rating:      userFromDb.Rating,
	}, nil
}

func (r *Repository) CheckUserData(ctx context.Context, name, phoneNumber, email string) error {
	existingUser := models.User{}

	err := r.db.Get(&existingUser, "SELECT name FROM users WHERE name = $1 AND deleted_at IS NULL", name)
	if err == nil {
		return nameAlreadyInUseErr
	}
	err = r.db.Get(&existingUser.Name, "SELECT id FROM users WHERE phone_number = $1 AND deleted_at IS NULL", phoneNumber)
	if err == nil {
		return phoneAlreadyInUseErr
	}
	err = r.db.Get(&existingUser, "SELECT id FROM users WHERE email = $1 AND deleted_at IS NULL", email)
	if err == nil {
		return emailAlreadyInUseErr
	}

	return nil
}

func (r *Repository) IsUserDeleted(ctx context.Context, userId int) bool {
	var deletedAt time.Time
	err := r.db.Get(&deletedAt, "SELECT deleted_at FROM users WHERE id = $1 AND deleted_at IS NULL ", userId)
	if err != nil {
		return false
	}
	return deletedAt != time.Unix(0, 0)
}
