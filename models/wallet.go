package models

type Wallet struct {
	Id       int            `json:"id"`
	Users    []WalletMember `json:"users"`
	Balance  float64        `json:"balance"`
	Owner    WalletMember   `json:"owner"`
	IsFamily bool           `json:"is_family"`
}

type Transaction struct {
	Id         int     `json:"id"`
	FromWallet Wallet  `json:"from_wallet"`
	ToWallet   Wallet  `json:"to_wallet"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}

type FamilyWallet struct {
	WalletId int  `json:"wallet_id"`
	UserId   int  `json:"user_id"`
	IsOwner  bool `json:"is_owner"`
}

type WalletMember struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Rating      string `json:"rating"`
}

type WalletFromDB struct {
	Id       int     `json:"id"`
	OwnerId  int     `json:"user_id"`
	Balance  float64 `json:"balance"`
	IsFamily bool    `json:"is_family"`
}

type WalletHistory struct {
	Transactions []Transaction `json:"transactions"`
}
