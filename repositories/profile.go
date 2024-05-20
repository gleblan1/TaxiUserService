package repositories

import (
	"context"
	"strconv"
	"strings"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type accountModel struct {
	id     int     `db:"id"`
	name   string  `db:"name"`
	phone  string  `db:"phone_number"`
	email  string  `db:"email"`
	rating float32 `db:"rating"`
}

func (r *Repository) GetAccountInfo(id int) (models.User, error) {
	var accountInfo accountModel
	var user models.User
	err := r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE id = $1", id).Scan(&accountInfo.id, &accountInfo.name, &accountInfo.phone, &accountInfo.email, &accountInfo.rating)
	if err != nil {
		return models.User{}, err
	}
	user = models.User{
		Id:          accountInfo.id,
		Name:        accountInfo.name,
		PhoneNumber: accountInfo.phone,
		Email:       accountInfo.email,
		Rating:      accountInfo.rating,
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
		keys, cursor, err = r.redis.Scan(ctx, cursor, "*", 10000).Result()
		if err != nil {
			return err
		}
		for _, key := range keys {
			val := strings.Split(key, ".")[0]
			if val == strconv.Itoa(id) {
				r.redis.Del(ctx, key)
			}
		}

		if cursor == 0 {
			break
		}
		keys, cursor, err = r.redis.Scan(ctx, cursor, "*", 10000).Result()
	}
	return nil
}
