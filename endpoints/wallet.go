package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

func GetWalletInfo(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.GetWalletInfoRequest)
		return UserService.GetWalletInfo(ctx, requestBody)
	}
}

func CashInWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.CashInWalletRequest)
		return UserService.CashInWallet(ctx, requestBody)
	}
}

func AddUserToWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.AddUserToWalletRequest)
		return UserService.AddUserToWallet(ctx, requestBody)
	}
}

func GetWalletTransactions(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.GetWalletTransactionsRequest)
		return UserService.GetWalletTransactions(ctx, requestBody)
	}
}

func ChooseWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.ChooseWalletRequest)

		return UserService.ChooseWallet(ctx, requestBody)
	}
}

func Pay(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.PayRequest)
		return UserService.Pay(ctx, requestBody)
	}
}

func CreateWallet(UserService services.UserService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestBody := request.(requests.CreateWalletRequest)
		return UserService.CreateWallet(ctx, requestBody)
	}
}
