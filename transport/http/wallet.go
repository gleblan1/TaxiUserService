package http

import (
	"net/http"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin"
)

type GetWalletInfoRequest struct {
	UserId int `json:"wallet_id"`
}

type CashInWalletRequest struct {
	WalletId int     `json:"wallet_id"`
	UserId   int     `json:"user_id"`
	Amount   float64 `json:"amount"`
}

type AddUserToWalletRequest struct {
	WalletId  int `json:"wallet_id"`
	UserId    int `json:"user_id"`
	UserToAdd int `json:"user_to_add"`
}

type GetWalletTransactionsRequest struct {
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

	requestBody := config.GetWalletInfoRequest{
		UserId: id,
	}

	result, err := h.e.GetWalletInfo(c, requestBody)
	if err != nil {

		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, result)
	return
}

func (h *Handler) CashInWallet(c *gin.Context) {
	var req CashInWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DefineResponse(c, http.StatusBadRequest, err)
		return
	}

	requestBody := config.CashInWalletRequest{
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

	requestBody := config.AddUserToWalletRequest{
		WalletId:  req.WalletId,
		UserId:    req.UserId,
		UserToAdd: req.UserToAdd,
	}
	if err, _ := h.e.AddUserToWallet(c, requestBody); err != nil {
		return
	}
	return
}

func (h *Handler) GetWalletTransactions(c *gin.Context) {
	var req GetWalletTransactionsRequest

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

	requestBody := config.GetWalletTransactionsRequest{
		UserId: id,
	}
	//TODO: сделать определение кода ошибки внутри функции дефайн респонс фффффффффффффффффффффффффффф
	response, err := h.e.GetWalletTransactions(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, response)
	return
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

	requestBody := config.ChooseWalletRequest{
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

	requestBody := config.PayRequest{
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

	requestBody := config.CreateWalletRequest{
		UserId:   id,
		IsFamily: req.IsFamily,
	}

	tokens, err := h.e.CreateWallet(c, requestBody)
	if err != nil {
		utils.DefineResponse(c, http.StatusUnauthorized, err)
		return
	}
	utils.DefineResponse(c, http.StatusOK, err, tokens)
	return
}
