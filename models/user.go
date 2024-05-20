package models

type User struct {
	Id          int
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Rating      float32
}

type UserInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Rating      string `json:"rating"`
}
