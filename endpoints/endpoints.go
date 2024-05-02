package endpoints

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
)

type Endpoints struct {
	SignUp        Endpoint
	SignIn        Endpoint
	SignOut       Endpoint
	RefreshTokens Endpoint

	GetAccountInfo Endpoint
	UpdateProfile  Endpoint
	DeleteProfile  Endpoint

	GetWalletInfo         Endpoint
	CashInWallet          Endpoint
	AddUserToWallet       Endpoint
	GetWalletTransactions Endpoint
	ChooseWallet          Endpoint
	Pay                   Endpoint
	CreateWallet          Endpoint

	RateTrip        Endpoint
	GetTripsHistory Endpoint
}

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func MakeEndpoints(UserService services.UserService) *Endpoints {
	return &Endpoints{
		SignUp:        SignUp(UserService),
		SignIn:        SignIn(UserService),
		SignOut:       SignOut(UserService),
		RefreshTokens: RefreshTokens(UserService),

		GetAccountInfo: GetAccountInfo(UserService),
		UpdateProfile:  UpdateProfile(UserService),
		DeleteProfile:  DeleteProfile(UserService),

		GetWalletInfo:         GetWalletInfo(UserService),
		CashInWallet:          CashInWallet(UserService),
		AddUserToWallet:       AddUserToWallet(UserService),
		GetWalletTransactions: GetWalletTransactions(UserService),
		ChooseWallet:          ChooseWallet(UserService),
		Pay:                   Pay(UserService),
		CreateWallet:          CreateWallet(UserService),

		RateTrip:        RateTrip(UserService),
		GetTripsHistory: GetTripsHistory(UserService),
	}
}
