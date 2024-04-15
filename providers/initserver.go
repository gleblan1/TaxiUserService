package providers

import (
	"fmt"
	handler "github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func run(port string, server *gin.Engine) error {
	err := server.Run(":" + port)
	if err != nil {
		return err
	}
	return nil
}

func InitServer() error {
	postgresDB, redisDB, err := InitDB()
	if err != nil {
		return fmt.Errorf("init postgres db err: %v", err)
	}
	repos := repositories.NewRepository(postgresDB, *redisDB)
	service := services.NewServices(repos)
	handler := handler.NewHandler(service)
	if err := run("8000", handler.InitRoutes()); err != nil {
		return fmt.Errorf("cannot run the server: %w", err)
	}

	return nil
}
