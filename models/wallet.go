package models

type Wallet struct {
	Id       int
	Users    []WalletMember
	Balance  int64
	Owner    WalletMember
	IsFamily bool
}

type Transaction struct {
	Id         int
	FromWallet Wallet
	ToWallet   Wallet
	Amount     int64
	Status     string
}

type FamilyWallet struct {
	WalletId int
	UserId   int
	IsOwner  bool
}

type WalletMember struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
	Rating      float32
}

type WalletFromDB struct {
	Id       int
	OwnerId  int
	Balance  int64
	IsFamily bool
}

type WalletHistory struct {
	Transactions []Transaction
}
