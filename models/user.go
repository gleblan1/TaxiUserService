package models

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required,min=4,max=20"`
	Email       string `json:"email" binding:"required,emailValid"`
	PhoneNumber string `json:"phone_number" binding:"required,phoneValid"`
	Password    string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=20"`
	Email       string `json:"email" binding:"required,emailValid"`
	PhoneNumber string `json:"phone_number" binding:"required,phoneValid"`
	Password    string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required,phoneValid"`
	Password    string `json:"password" binding:"required,min=8"`
}
