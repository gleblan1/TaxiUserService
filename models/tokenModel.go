package models

type JwtClaims struct {
	Audience  string  `json:"aud,omitempty"`
	ExpiresAt float64 `json:"exp,omitempty"`
	IssuedAt  float64 `json:"iat,omitempty"`
	Session   string  `'json:"jti,omitempty"`
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}
