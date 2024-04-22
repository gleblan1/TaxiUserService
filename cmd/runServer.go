package main

import (
	"context"
	"fmt"

	handler "github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/middleware"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/providers"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, stop context.CancelFunc) error {
	g, gCtx := errgroup.WithContext(ctx)
	postgresDB, redisDB, err := providers.InitDB()
	if err != nil {
		return fmt.Errorf("init postgres db err: %v", err)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("emailValid", utils.EmailValid)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("phoneValid", utils.PhoneValid)
		if err != nil {
			return err
		}
	}
	repos := repositories.NewRepository(
		repositories.WithPostgresRepository(postgresDB),
		repositories.WithRedisClient(*redisDB),
	)
	service := services.NewService(
		services.WithAuthRepo(repos),
	)
	handlers := handler.NewHandler(
		handler.WithAuthService(service),
	)
	middlewares := middleware.NewMiddleware(
		middleware.WithAuthMiddleware(service),
	)
	router := handler.NewRouter(
		handler.WithHandler(handlers),
		handler.WithMiddleware(middlewares),
	)
	g.Go(func() error {
		if err := providers.InitServer("8000", router.InitRoutes()); err != nil {
			return fmt.Errorf("cannot run the server: %w", err)
		}
		return nil
	})

	select {
	case <-gCtx.Done():
		stop()
		return g.Wait()
	case <-ctx.Done():
		fmt.Println(" Exited")
		return nil
	}
}
