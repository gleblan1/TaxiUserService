package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/shopspring/decimal"
)

type transactionFromDB struct {
	id         int
	walletFrom int
	walletTo   int
	amount     int64
	status     string
}

func (r *Repository) GetWalletInfo(ctx context.Context, userId int) (models.Wallet, error) {
	var walletFromDB models.WalletFromDB
	currentWalletId, err := r.GetCurrentWalletId(ctx, userId)
	if err != nil {
		return models.Wallet{}, err
	}

	err = r.db.QueryRow("SELECT id, user_id, balance, is_family FROM wallets WHERE id = $1", currentWalletId).Scan(&walletFromDB.Id, &walletFromDB.OwnerId, &walletFromDB.Balance, &walletFromDB.IsFamily)
	if err != nil {
		return models.Wallet{}, errors.New("choose wallet please" + err.Error())
	}

	var owner models.WalletMember
	var members []models.WalletMember

	err = r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE current_wallet_id = $1", currentWalletId).Scan(&owner.Id, &owner.Name, &owner.PhoneNumber, &owner.Email, &owner.Rating)
	if err != nil {
		fmt.Println(err)
		return models.Wallet{}, errors.New("choose wallet please")
	}

	rows, err := r.db.Query("SELECT user_id FROM family_wallets WHERE wallet_id = $1", walletFromDB.Id)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var member models.WalletMember

		err = rows.Scan(&member.Id)
		if err != nil {
			return models.Wallet{}, err
		}
		_ = r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE id = $1", member.Id).Scan(&member.Id, &member.Name, &member.PhoneNumber, &member.Email, &member.Rating)
		members = append(members, member)
	}

	wallet := models.Wallet{
		Id:       walletFromDB.Id,
		Balance:  walletFromDB.Balance,
		IsFamily: walletFromDB.IsFamily,
		Owner:    owner,
		Users:    append(members, owner),
	}

	return wallet, nil
}

func (r *Repository) CashInWallet(ctx context.Context, walletID int, amount int64) (models.Wallet, error) {
	var currentBalance int64
	err := r.db.QueryRow("SELECT balance FROM wallets WHERE id = $1", walletID).Scan(&currentBalance)
	if err != nil {
		fmt.Println(err)
		return models.Wallet{}, err
	}

	finalBalance := currentBalance + amount
	_, err = r.db.Exec("UPDATE wallets SET balance = $1 WHERE id = $2", finalBalance, walletID)
	if err != nil {
		return models.Wallet{}, errors.New("wallet is not exists")
	}

	return models.Wallet{}, nil
}

func (r *Repository) AddUserToWallet(ctx context.Context, walletID, userToAdd, userId int) (models.Wallet, error) {
	wallet := models.Wallet{}
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM family_wallets WHERE wallet_id = $1 AND user_id = $2", walletID, userToAdd).Scan(&count)
	if err != nil {
		return models.Wallet{}, err
	}

	if count > 0 {
		return models.Wallet{}, errors.New("user already added in your family account")
	}

	_, err = r.db.Exec("INSERT INTO family_wallets (wallet_id, user_id, is_owner) VALUES ($1, $2, $3)", walletID, userToAdd, false)
	if err != nil {
		return models.Wallet{}, err
	}

	_, err = r.db.Exec("UPDATE wallets SET is_family=true WHERE id=$1", walletID)
	if err != nil {
		return models.Wallet{}, err
	}

	wallet, err = r.GetWalletById(ctx, walletID)
	if err != nil {
		return models.Wallet{}, err
	}

	return wallet, nil
}

func (r *Repository) GetWalletById(ctx context.Context, walletId int) (models.Wallet, error) {
	var wallet models.Wallet
	var owner models.WalletMember
	var members []models.WalletMember
	ownerId, err := r.GetOwnerOfWallet(ctx, walletId)
	err = r.db.QueryRow("SELECT id, is_family, balance FROM wallets WHERE id = $1", walletId).Scan(&wallet.Id, &wallet.IsFamily, &wallet.Balance)
	if err != nil {
		return models.Wallet{}, err
	}

	err = r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE id = $1", ownerId).Scan(&owner.Id, &owner.Name, &owner.PhoneNumber, &owner.Email, &owner.Rating)
	if err != nil {
		return models.Wallet{}, err
	}

	rows, err := r.db.Query("SELECT user_id FROM family_wallets WHERE wallet_id=$1", wallet.Id)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return models.Wallet{}, err
	}

	for rows.Next() {
		var member models.WalletMember

		err = rows.Scan(&member.Id)
		if err != nil {
			return models.Wallet{}, err
		}
		_ = r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE id = $1", member.Id).Scan(&member.Id, &member.Name, &member.PhoneNumber, &member.Email, &member.Rating)
		members = append(members, member)
	}
	wallet = models.Wallet{
		Id:       wallet.Id,
		Users:    append(members, owner),
		Balance:  wallet.Balance,
		Owner:    owner,
		IsFamily: wallet.IsFamily,
	}
	return wallet, nil
}

func (r *Repository) GetWalletTransactions(ctx context.Context, walletId int) (models.WalletHistory, error) {
	var transactions []models.Transaction
	var walletHistory models.WalletHistory
	var transactionFromQuery transactionFromDB
	rows, err := r.db.Query("SELECT id FROM transactions WHERE from_wallet = $1", walletId)
	if err != nil {
		return walletHistory, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var transaction models.Transaction

		err = rows.Scan(&transaction.Id)
		if err != nil {
			return walletHistory, err
		}
		_ = r.db.QueryRow("SELECT id, from_wallet, to_wallet, amount, status FROM transactions WHERE id = $1", transaction.Id).Scan(&transactionFromQuery.id, &transactionFromQuery.walletFrom, &transactionFromQuery.walletTo, &transactionFromQuery.amount, &transactionFromQuery.status)
		var walletFrom models.Wallet
		var walletTo models.Wallet

		walletFrom, err = r.GetWalletById(ctx, transactionFromQuery.walletFrom)
		if err != nil {
			return walletHistory, err
		}
		walletTo, err = r.GetWalletById(ctx, transactionFromQuery.walletTo)
		if err != nil {
			return walletHistory, err
		}
		transaction = models.Transaction{
			Id:         transactionFromQuery.id,
			FromWallet: walletFrom,
			ToWallet:   walletTo,
			Amount:     transactionFromQuery.amount,
			Status:     transactionFromQuery.status,
		}
		transactions = append(transactions, transaction)
	}

	walletHistory = models.WalletHistory{
		Transactions: transactions,
	}
	return walletHistory, nil
}

func (r *Repository) GetOwnerOfWallet(ctx context.Context, walletID int) (int, error) {
	var userOwnerId int
	err := r.db.QueryRow("SELECT user_id FROM wallets WHERE id =$1", walletID).Scan(&userOwnerId)
	if err != nil {
		return 0, err
	}
	return userOwnerId, nil
}

func (r *Repository) ChooseWallet(ctx context.Context, walletID, userId int) (models.Wallet, error) {
	_, err := r.db.Exec("UPDATE users SET current_wallet_id = $1 WHERE id = $2", walletID, userId)
	return models.Wallet{}, err
}

func (r *Repository) GetCurrentWalletId(ctx context.Context, userId int) (int, error) {
	var id int
	err := r.db.QueryRow("SELECT current_wallet_id FROM users WHERE id = $1", userId).Scan(&id)
	if err != nil {
		return 0, errors.New("choose wallet")
	}
	return id, nil
}

func (r *Repository) GetBalance(ctx context.Context, walletId int) decimal.Decimal {
	var balance float64

	err := r.db.QueryRow("SELECT balance FROM wallets WHERE id = $1", walletId).Scan(&balance)
	if err != nil {
		return decimal.NewFromFloat(0.0)
	}

	return decimal.NewFromFloatWithExponent(balance, -2)
}

func (r *Repository) Pay(ctx context.Context, walletID, toWalletId int, amount int64) (models.Wallet, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return models.Wallet{}, err
	}
	defer (func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			_, err = tx.ExecContext(ctx, "UPDATE transactions SET status = 'failed'")
			if err != nil {
				return
			}
		}
	})()

	var id int64
	err = tx.QueryRowContext(ctx, "INSERT INTO transactions (from_wallet, to_wallet, amount, status) VALUES ($1, $2, $3, $4) RETURNING id", walletID, toWalletId, amount, "started").Scan(&id)
	if err != nil {
		return models.Wallet{}, err
	}
	_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance = balance - $1 WHERE id = $2", amount, walletID)
	if err != nil {
		return models.Wallet{}, err
	}

	_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance = balance + $1 WHERE id = $2", amount, toWalletId)
	if err != nil {
		return models.Wallet{}, err
	}

	_, err = tx.ExecContext(ctx, "UPDATE transactions SET status = 'success' WHERE id=$1", id)
	if err != nil {
		return models.Wallet{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.Wallet{}, err
	}

	return models.Wallet{}, nil
}

func (r *Repository) CreateWallet(ctx context.Context, userId int, isFamily bool) (models.Wallet, error) {
	var wallet models.Wallet
	var user models.WalletMember
	err := r.db.QueryRow("SELECT id, name, phone_number, email, rating FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Email, &user.Rating)
	if err != nil {
		return models.Wallet{}, err
	}

	stmt, err := r.db.Prepare("INSERT INTO wallets(user_id, is_family, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id")
	if err != nil {
		return wallet, fmt.Errorf("error preparing SQL statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
		}
	}(stmt)

	err = stmt.QueryRow(userId, isFamily).Scan(&wallet.Id)
	if err != nil {
		return wallet, fmt.Errorf("error executing SQL statement: %w", err)
	}

	wallet = models.Wallet{
		Id:       wallet.Id,
		Users:    append(wallet.Users, user),
		Owner:    user,
		IsFamily: isFamily,
	}

	return wallet, nil

}
