package services

import (
	"context"
	"errors"

	"github.com/shopspring/decimal"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

func (s *Service) GetWalletInfo(ctx context.Context, req requests.GetWalletInfoRequest) (models.Wallet, error) {
	var walletInfo models.Wallet
	wallet, err := s.repo.GetWalletInfo(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}
	walletInfo = models.Wallet{
		Id:       wallet.Id,
		Users:    wallet.Users,
		Balance:  wallet.Balance,
		Owner:    wallet.Owner,
		IsFamily: wallet.IsFamily,
	}
	return walletInfo, nil
}

func (s *Service) CashInWallet(ctx context.Context, req requests.CashInWalletRequest) (models.Wallet, error) {
	amount := utils.ConvertFloatToInt(req.Amount)
	return s.repo.CashInWallet(ctx, req.WalletId, amount)
}

func (s *Service) CheckIsOwner(ctx context.Context, userId, walletId int) bool {
	realOwner, err := s.repo.GetOwnerOfWallet(ctx, walletId)
	if err != nil {
		return false
	}
	if realOwner == userId {
		return true
	}
	return false
}

func (s *Service) AddUserToWallet(ctx context.Context, req requests.AddUserToWalletRequest) (models.Wallet, error) {
	walletID, err := s.repo.GetCurrentWalletId(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}
	if s.CheckIsOwner(ctx, req.UserToAdd, req.UserId) {
		return models.Wallet{}, errors.New("you cannot add yourself to your wallet")
	}
	return s.repo.AddUserToWallet(ctx, walletID, req.UserToAdd, req.UserId)
}

func (s *Service) GetWalletTransactions(ctx context.Context, req requests.GetWalletTransactionsRequest) (models.WalletHistory, error) {
	currentWalletId, err := s.repo.GetCurrentWalletId(ctx, req.UserId)
	if err != nil {
		return models.WalletHistory{}, err
	}
	return s.repo.GetWalletTransactions(ctx, currentWalletId)
}

func (s *Service) ChooseWallet(ctx context.Context, req requests.ChooseWalletRequest) (models.Wallet, error) {
	userOwnerId, err := s.repo.GetOwnerOfWallet(ctx, req.WalletId)
	if err != nil {
		return models.Wallet{}, err
	}
	if userOwnerId == req.UserId {
		_, err := s.repo.ChooseWallet(ctx, req.WalletId, req.UserId)
		if err != nil {
			return models.Wallet{}, errors.New("wallet is not exists")
		}
		return models.Wallet{}, nil
	}

	return models.Wallet{}, errors.New("it is not your wallet")

}

func (s *Service) Pay(ctx context.Context, req requests.PayRequest) (models.Wallet, error) {
	currentWalletId, err := s.repo.GetCurrentWalletId(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}

	balance := s.repo.GetBalance(ctx, currentWalletId)
	amount := decimal.NewFromFloatWithExponent(req.Amount, -2)

	finalValue := balance.Sub(amount)
	if finalValue.LessThan(decimal.NewFromFloatWithExponent(0, -2)) {
		return models.Wallet{}, errors.New("not enough balance")
	} else if req.Amount <= 0 {
		return models.Wallet{}, errors.New("incorrect value")
	}
	return s.repo.Pay(ctx, currentWalletId, req.ToWalletId, utils.ConvertDecimalToInt(amount))
}

func (s *Service) CreateWallet(ctx context.Context, req requests.CreateWalletRequest) (models.Wallet, error) {
	return s.repo.CreateWallet(ctx, req.UserId, req.IsFamily)
}
