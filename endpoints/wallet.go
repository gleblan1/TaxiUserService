package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func GetWalletInfo(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.GetWalletInfoRequest)
		return UserService.GetWalletInfo(ctx, requestBody)
	}
}

func CashInWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.CashInWalletRequest)
		return UserService.CashInWallet(ctx, requestBody)
	}
}

func AddUserToWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.AddUserToWalletRequest)
		return UserService.AddUserToWallet(ctx, requestBody)
	}
}

func GetWalletTransactions(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.GetWalletTransactionsRequest)
		return UserService.GetWalletTransactions(ctx, requestBody)
	}
}

func ChooseWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.ChooseWalletRequest)

		return UserService.ChooseWallet(ctx, requestBody)
	}
}

func Pay(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.PayRequest)
		return UserService.Pay(ctx, requestBody)
	}
}

func CreateWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(config.CreateWalletRequest)
		return UserService.CreateWallet(ctx, requestBody)
	}
}
