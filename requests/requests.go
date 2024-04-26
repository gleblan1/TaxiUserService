package requests

//auth

type RegisterRequest struct {
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type SignInRequest struct {
	PhoneNumber string
	Password    string
}

type LogoutRequest struct {
	SessionId int
	UserId    int
}

type RefreshTokensRequest struct {
	RefreshToken string
}

type PatchRequest struct {
	Name        string
	PhoneNumber string
	Email       string
}

//profile

type GetAccountInfoRequest struct {
	Id int
}

type UpdateProfileRequest struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
}

type DeleteProfileRequest struct {
	Id int
}

//wallets

type GetWalletInfoRequest struct {
	UserId int
}

type CashInWalletRequest struct {
	WalletId int
	Amount   float64
}

type AddUserToWalletRequest struct {
	WalletId  int
	UserId    int
	UserToAdd int
}

type GetWalletTransactionsRequest struct {
	UserId int
}

type ChooseWalletRequest struct {
	WalletId int
	UserId   int
}

type PayRequest struct {
	UserId     int
	ToWalletId int
	Amount     float64
}

type CreateWalletRequest struct {
	UserId   int
	IsFamily bool
}
