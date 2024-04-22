package models

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Rating      string `json:"rating"`
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

type UserInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Rating      string `json:"rating"`
}

type PatchRequest struct {
	Name        string `json:"name" binding:"omitempty,min=4,max=20"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,phoneValid"`
	Email       string `json:"email" binding:"omitempty,emailValid"`
}
