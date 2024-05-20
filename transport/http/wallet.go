package http

import (
	"net/http"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type GetWalletInfoRequest struct {
	UserId int `json:"wallet_id"`
}

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id         int                   `json:"id"`
	FromWallet GetWalletInfoResponse `json:"from_wallet"`
	ToWallet   GetWalletInfoResponse `json:"to_wallet"`
	Amount     float64               `json:"amount"`
	Status     string                `json:"status"`
}

type GetWalletInfoResponse struct {
	Id       int            `json:"id"`
	Users    []WalletMember `json:"users"`
	Balance  float64        `json:"balance"`
	Owner    WalletMember   `json:"owner"`
	IsFamily bool           `json:"is_family"`
}

type WalletMember struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Rating      float32 `json:"rating"`
}

type CashInWalletRequest struct {
	WalletId int     `json:"wallet_id"`
	UserId   int     `json:"user_id"`
	Amount   float64 `json:"amount"`
}

type AddUserToWalletRequest struct {
	WalletId  int `json:"wallet_id"`
	UserToAdd int `json:"user_to_add"`
}

type ChooseWalletRequest struct {
	WalletId int `json:"wallet_id"`
	UserId   int `json:"user_id"`
}

type PayRequest struct {
	ToWalletId int     `json:"to_wallet_id"`
	Amount     float64 `json:"amount"`
}

type CreateWalletRequest struct {
	UserId   int  `json:"user_id"`
	IsFamily bool `json:"is_family"`
}

func (h *Handler) GetWalletInfo(c *gin.Context) {
	var req GetWalletInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.GetWalletInfoRequest{
		UserId: id,
	}

	wallet, err := h.e.GetWalletInfo(c, requestBody)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	result := generateWalletResponse(id, wallet)

	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}

func (h *Handler) CashInWallet(c *gin.Context) {
	var req CashInWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.CashInWalletRequest{
		WalletId: req.WalletId,
		Amount:   req.Amount,
	}

	_, err := h.e.CashInWallet(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, "you successfully cashed in your wallet")
	return
}

func (h *Handler) AddUserToWallet(c *gin.Context) {
	var req AddUserToWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.AddUserToWalletRequest{
		UserId:    id,
		UserToAdd: req.UserToAdd,
	}
	wallet, err := h.e.AddUserToWallet(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	result := generateWalletResponse(id, wallet)
	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}

func (h *Handler) GetWalletTransactions(c *gin.Context) {

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.GetWalletTransactionsRequest{
		UserId: id,
	}

	transactions, err := h.e.GetWalletTransactions(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}

	result := TransactionsResponse{}
	for i, v := range transactions.(models.WalletHistory).Transactions {
		result.Transactions = append(result.Transactions, Transaction{})
		result.Transactions[i].Id = v.Id
		result.Transactions[i].FromWallet = generateWalletResponse(v.FromWallet.Id, v.FromWallet)
		result.Transactions[i].ToWallet = generateWalletResponse(v.ToWallet.Id, v.ToWallet)
		result.Transactions[i].Amount = utils.ConvertIntToFloat(v.Amount)
		result.Transactions[i].Status = v.Status

	}
	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}
func generateWalletResponse(id int, wallet interface{}) GetWalletInfoResponse {
	result := GetWalletInfoResponse{
		Id:       id,
		Users:    make([]WalletMember, len(wallet.(models.Wallet).Users)),
		Balance:  utils.ConvertIntToFloat(wallet.(models.Wallet).Balance) / 100,
		IsFamily: wallet.(models.Wallet).IsFamily,
	}

	for i, user := range wallet.(models.Wallet).Users {
		result.Users[i].Id = user.Id
		result.Users[i].Name = user.Name
		result.Users[i].PhoneNumber = user.PhoneNumber
		result.Users[i].Email = user.Email
		result.Users[i].Rating = user.Rating
	}

	result.Owner.Id = wallet.(models.Wallet).Owner.Id
	result.Owner.Name = wallet.(models.Wallet).Owner.Name
	result.Owner.PhoneNumber = wallet.(models.Wallet).Owner.PhoneNumber
	result.Owner.Email = wallet.(models.Wallet).Owner.Email
	result.Owner.Rating = wallet.(models.Wallet).Owner.Rating

	return result
}
func (h *Handler) ChooseWallet(c *gin.Context) {
	var req ChooseWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.ChooseWalletRequest{
		WalletId: req.WalletId,
		UserId:   id,
	}

	_, err = h.e.ChooseWallet(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, "wallet chosen")
	return
}

func (h *Handler) Pay(c *gin.Context) {
	var req PayRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.PayRequest{
		UserId:     id,
		ToWalletId: req.ToWalletId,
		Amount:     req.Amount,
	}

	_, err = h.e.Pay(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, "success")
	return
}

func (h *Handler) CreateWallet(c *gin.Context) {
	var req CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	claims, err := utils.ExtractClaims(utils.GetTokenFromHeader(c))
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(claims.Audience)
	if err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := requests.CreateWalletRequest{
		UserId:   id,
		IsFamily: req.IsFamily,
	}

	wallet, err := h.e.CreateWallet(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}

	result := generateWalletResponse(id, wallet)

	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}
