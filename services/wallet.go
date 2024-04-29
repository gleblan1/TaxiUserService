package services

import (
	"context"
	"errors"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/shopspring/decimal"
)

type ResponseWalletInfo struct {
	Id       int
	Users    []models.WalletMember
	Balance  float64
	Owner    models.WalletMember
	IsFamily bool
}

func (s *Service) GetWalletInfo(ctx context.Context, req config.GetWalletInfoRequest) (ResponseWalletInfo, error) {
	var walletInfo ResponseWalletInfo
	wallet, err := s.repo.GetWalletInfo(ctx, req.UserId)
	if err != nil {
		return ResponseWalletInfo{}, err
	}
	walletInfo = ResponseWalletInfo{
		Id:       wallet.Id,
		Users:    wallet.Users,
		Balance:  s.ConvertIntToFloat(wallet.Balance),
		Owner:    wallet.Owner,
		IsFamily: wallet.IsFamily,
	}
	return walletInfo, nil
}

func (s *Service) CashInWallet(ctx context.Context, req config.CashInWalletRequest) (models.Wallet, error) {
	amount := s.ConvertFloatToInt(req.Amount)
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
	currentWalletId, err := s.repo.GetCurrentWalletId(ctx, req.UserId)
	if err != nil {
		return models.WalletHistory{}, err
	}
	return s.repo.GetWalletTransactions(ctx, currentWalletId)
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

func (s *Service) ConvertDecimalToInt(value decimal.Decimal) int64 {
	return value.Mul(decimal.NewFromFloat(100)).IntPart()
}

func (s *Service) ConvertFloatToInt(value float64) int64 {
	return decimal.NewFromFloatWithExponent(value, -2).Mul(decimal.NewFromFloat(100)).IntPart()
}

func (s *Service) ConvertIntToFloat(value int64) float64 {
	result, _ := decimal.NewFromFloatWithExponent(float64(value), -2).Div(decimal.NewFromFloat(100)).Float64()
	return result
}

func (s *Service) Pay(ctx context.Context, req config.PayRequest) (models.Wallet, error) {
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
	return s.repo.Pay(ctx, currentWalletId, req.ToWalletId, s.ConvertDecimalToInt(amount))
}

func (s *Service) CreateWallet(ctx context.Context, req config.CreateWalletRequest) (models.Wallet, error) {
	return s.repo.CreateWallet(ctx, req.UserId, req.IsFamily)
}
