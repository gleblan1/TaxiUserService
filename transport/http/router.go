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
	auth.POST("/login", r.handler.LogIn)
	auth.POST("/sign-up", r.handler.SignUp)
	auth.POST("/refresh", r.handler.Refresh)
	auth.POST("/logout", r.authMiddleware.ValidateToken(), r.handler.LogOut)
	return router
}
