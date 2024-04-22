package handler

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	authMiddleware IAuthMiddleware
	handler        Handler
}

func NewRouter(authMiddleware IAuthMiddleware, handler Handler) *Router {
	return &Router{authMiddleware: authMiddleware, handler: handler}
}

type IAuthMiddleware interface {
	ValidateToken() gin.HandlerFunc
}

func (r *Router) InitRoutes() *gin.Engine {

	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/login", r.handler.LogIn)
	auth.POST("/sign-up", r.handler.SignUp)
	auth.POST("/refresh", r.handler.Refresh)
	auth.POST("/logout", r.authMiddleware.ValidateToken(), r.handler.LogOut)
	return router
}
