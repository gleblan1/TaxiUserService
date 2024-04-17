package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func EmailValid(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func PhoneValid(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	phoneRegex := regexp.MustCompile(`^\+\d{12}$`)
	return phoneRegex.MatchString(phone)
}
