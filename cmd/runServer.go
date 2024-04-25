package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/middleware"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/providers"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/transport/http"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

func Run(ctx context.Context, stop context.CancelFunc) error {
	g, gCtx := errgroup.WithContext(ctx)
	psql, err := providers.InitPostgres()
	if err != nil {
		return err
	}
	redis, err := providers.InitRedis(ctx)
	if err != nil {
		return err
	}
	DB := providers.NewDB(
		providers.WithPsql(psql),
		providers.WithRedis(redis),
	)
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
		repositories.WithPostgresRepository(DB.Psql),
		repositories.WithRedisClient(*DB.Redis),
	)

	service := services.NewService(
		services.WithAuthRepo(repos),
	)

	endpoint := endpoints.MakeEndpoints(service)

	handlers := http.NewHandler(
		http.WithAuthService(endpoint),
	)

	middlewares := middleware.NewMiddleware(
		middleware.WithAuthMiddleware(service),
	)

	router := http.NewRouter(
		http.WithHandler(handlers),
		http.WithMiddleware(middlewares),
	)
	g.Go(func() error {
		server := providers.NewServer(
			providers.WithPort(utils.ReadValue("PORT")),
			providers.WithServer(router.InitRoutes()),
		)
		if err := providers.InitServer(server.Port, server.Server); err != nil {
			return fmt.Errorf("cannot run the server: %w", err)
		}
		return nil
	})

	select {
	case <-gCtx.Done():
		stop()
		fmt.Println(" Exited")
		return nil
	}
}
