package requests

type RegisterRequest struct {
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LogoutRequest struct {
	SessionId int
	UserId    int
}

type RefreshRequestBody struct {
	RefreshToken string
}
