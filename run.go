package run

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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
	db, err := sqlx.Open("postgres", utils.DbConnectionString())
	if err != nil {
		return fmt.Errorf("db connection error: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("db ping error: %w", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis ping error: %w", err)
	}

	psql := repositories.NewPostgres(
		repositories.WithSqlxDB(db),
	)

	redisStore := repositories.NewRedis(
		repositories.WithRedisClient(client),
	)

	repos := repositories.NewRepository(
		repositories.WithPostgresRepository(*psql),
		repositories.WithRedis(*redisStore),
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
			providers.WithServer(router.NewRoutes()),
		)
		err := server.Server.Run(":" + server.Port)
		if err != nil {
			return err
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
