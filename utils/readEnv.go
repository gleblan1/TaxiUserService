package utils

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func ReadValue(envString string) string {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.AutomaticEnv()

	value := viper.GetString(envString)
	return value
}
