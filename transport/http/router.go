package http

import (
	"github.com/gin-gonic/gin"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/middleware"
)

type Router struct {
	authMiddleware *middleware.Middleware
	handler        *Handler
}

type RouterOptions func(*Router)

func NewRouter(options ...RouterOptions) *Router {
	router := &Router{}
	for _, option := range options {
		option(router)
	}
	return router
}

func WithMiddleware(middleware *middleware.Middleware) RouterOptions {
	return func(r *Router) {
		r.authMiddleware = middleware
	}
}

func WithHandler(h *Handler) RouterOptions {
	return func(r *Router) {
		r.handler = h
	}
}

func (r *Router) NewRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/sign-in", r.handler.SignIn)
	auth.POST("/sign-up", r.handler.SignUp)
	auth.POST("/refresh", r.handler.RefreshTokens)
	auth.POST("/logout", r.authMiddleware.ValidateToken(), r.handler.SignOut)

	profile := router.Group("/profile")
	profile.Use(r.authMiddleware.ValidateToken())
	profile.PATCH("/", r.handler.UpdateProfile)
	profile.DELETE("/", r.handler.DeleteProfile)
	profile.GET("/", r.handler.GetAccountInfo)

	wallet := router.Group("/wallet")
	wallet.Use(r.authMiddleware.ValidateToken())
	wallet.GET("/", r.handler.GetWalletInfo)
	wallet.GET("/transactions", r.handler.GetWalletTransactions)
	wallet.PATCH("/balance", r.handler.CashInWallet)
	wallet.POST("/users", r.handler.AddUserToWallet)
	wallet.POST("/payment", r.handler.Pay)
	wallet.PATCH("/", r.handler.ChooseWallet)
	wallet.POST("/", r.handler.CreateWallet)
	return router
}
