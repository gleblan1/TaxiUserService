package services

import (
	"context"
	"errors"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

func (s *Service) GetWalletInfo(ctx context.Context, req config.GetWalletInfoRequest) (models.Wallet, error) {
	wallet, err := s.repo.GetWalletInfo(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (s *Service) CashInWallet(ctx context.Context, req config.CashInWalletRequest) (models.Wallet, error) {
	return s.repo.CashInWallet(ctx, req.WalletId, req.Amount)
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

func (s *Service) AddUserToWallet(ctx context.Context, req config.AddUserToWalletRequest) (models.Wallet, error) {
	walletID, err := s.repo.GetCurrentWalletId(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}
	if s.CheckIsOwner(ctx, req.UserToAdd, req.UserId) {
		return models.Wallet{}, errors.New("you cannot add yourself to your wallet")
	}
	return s.repo.AddUserToWallet(ctx, walletID, req.UserToAdd, req.UserId)
}

func (s *Service) GetWalletTransactions(ctx context.Context, req config.GetWalletTransactionsRequest) (models.WalletHistory, error) {
	return s.repo.GetWalletTransactions(ctx, req.UserId)
}

func (s *Service) ChooseWallet(ctx context.Context, req config.ChooseWalletRequest) (models.Wallet, error) {
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

func (s *Service) Pay(ctx context.Context, req config.PayRequest) (models.Wallet, error) {
	if err := s.repo.CheckBalance(ctx, req.Amount, req.UserId); err != nil {
		return models.Wallet{}, err
	}
	return s.repo.Pay(ctx, req.UserId, req.ToWalletId, req.Amount)
}

func (s *Service) CreateWallet(ctx context.Context, req config.CreateWalletRequest) (models.Wallet, error) {
	return s.repo.CreateWallet(ctx, req.UserId, req.IsFamily)
}
