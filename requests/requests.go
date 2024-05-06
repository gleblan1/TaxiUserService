package requests

type RegisterRequest struct {
	Name        string
	Email       string
	PhoneNumber string
	Password    string
}

type SignInRequest struct {
	PhoneNumber string
	Password    string
}

type LogoutRequest struct {
	SessionId int
	UserId    int
}

type RefreshTokensRequest struct {
	RefreshToken string
}

type PatchRequest struct {
	Name        string
	PhoneNumber string
	Email       string
}

type GetAccountInfoRequest struct {
	Id int
}

type UpdateProfileRequest struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
}

type DeleteProfileRequest struct {
	Id int
}
