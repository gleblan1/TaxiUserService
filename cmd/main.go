package main

import (
	"fmt"
	handler "github.com/GO-Trainee/GlebL-innotaxi-userservice/endpoints"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/middleware"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/providers"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/repositories"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/services"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
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
	repos := repositories.NewRepository(postgresDB, *redisDB)
	service := services.NewServices(repos)
	handlers := handler.NewHandler(service)
	middlewares := middleware.NewAuthMiddleware(*service)
	router := handler.NewRouter(middlewares, *handlers)
	if err := providers.InitServer("8000", router.InitRoutes()); err != nil {
		return fmt.Errorf("cannot run the server: %w", err)
	}

	return nil
}
