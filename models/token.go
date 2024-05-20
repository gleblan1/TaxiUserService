package models

type JwtClaims struct {
	Audience  string
	ExpiresAt float64
	IssuedAt  float64
	Session   string
}

type JwtToken struct {
	AccessToken  string
	RefreshToken string
}
