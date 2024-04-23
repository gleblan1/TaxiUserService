package repositories

import (
	"context"
	"strconv"
	"strings"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type Profile interface {
	GetAccountInfo(id int) (models.UserInfo, error)
	UpdateUsername(name string, id int) error
	UpdatePhoneNumber(phoneNumber string, id int) error
	UpdateEmail(email string, id int) error
	UpdateDeletedStatus(id int) error
	DeleteAllSessions(ctx context.Context, id int) error
}

func (r *Repository) GetAccountInfo(id int) (models.UserInfo, error) {
	var user models.UserInfo
	err := r.db.QueryRow("SELECT name, phone_number, email, rating FROM users WHERE id = $1", id).Scan(&user.Name, &user.PhoneNumber, &user.Email, &user.Rating)
	if err != nil {
		return models.UserInfo{}, err
	}
	return user, nil
}

func (r *Repository) UpdateUsername(name string, id int) error {
	_, err := r.db.Exec("UPDATE users SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdatePhoneNumber(phoneNumber string, id int) error {
	_, err := r.db.Exec("UPDATE users SET phone_number = $1 WHERE id = $2", phoneNumber, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateEmail(email string, id int) error {
	_, err := r.db.Exec("UPDATE users SET email = $1 WHERE id = $2", email, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateDeletedStatus(id int) error {
	_, err := r.db.Exec("UPDATE users SET deleted_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteAllSessions(ctx context.Context, id int) error {
	var cursor uint64 = 0
	var keys []string
	for {
		var err error
		keys, cursor, err = r.client.Scan(ctx, cursor, "*", 10000).Result()
		if err != nil {
			return err
		}
		for _, key := range keys {
			val := strings.Split(key, ".")[0]
			if val == strconv.Itoa(id) {
				r.client.Del(ctx, key)
			}
		}

		if cursor == 0 {
			break
		}
		keys, cursor, err = r.client.Scan(ctx, cursor, "*", 10000).Result()
	}
	return nil
}
