package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type Wallet interface {
	GetWalletInfo(ctx context.Context, req config.GetWalletInfoRequest) (models.Wallet, error)
	CashInWallet(ctx context.Context, req config.CashInWalletRequest) (models.Wallet, error)
	AddUserToWallet(ctx context.Context, req config.AddUserToWalletRequest) (models.Wallet, error)
	GetWalletTransactions(ctx context.Context, req config.GetWalletTransactionsRequest) (models.WalletHistory, error)
	ChooseWallet(ctx context.Context, req config.ChooseWalletRequest) (models.Wallet, error)
	Pay(ctx context.Context, req config.PayRequest) (models.Wallet, error)
	CreateWallet(ctx context.Context, req config.CreateWalletRequest) (models.Wallet, error)
}

func (s *Service) GetWalletInfo(ctx context.Context, req config.GetWalletInfoRequest) (models.Wallet, error) {
	wallet, err := s.repo.GetWalletInfo(ctx, req.UserId)
	if err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (s *Service) CashInWallet(ctx context.Context, req config.CashInWalletRequest) (models.Wallet, error) {
	return s.repo.CashInWallet(ctx, req.WalletId, req.UserId, req.Amount)
}

func (s *Service) AddUserToWallet(ctx context.Context, req config.AddUserToWalletRequest) (models.Wallet, error) {
	return s.repo.AddUserToWallet(ctx, req.UserToAdd, req.UserId)
}

func (s *Service) GetWalletTransactions(ctx context.Context, req config.GetWalletTransactionsRequest) (models.WalletHistory, error) {
	return s.repo.GetWalletTransactions(ctx, req.UserId)
}

func (s *Service) ChooseWallet(ctx context.Context, req config.ChooseWalletRequest) (models.Wallet, error) {
	return s.repo.ChooseWallet(ctx, req.WalletId, req.UserId)
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
