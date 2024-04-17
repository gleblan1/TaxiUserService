package models

import (
	"time"
)

type JwtClaims struct {
	Audience  string  `json:"aud,omitempty"`
	ExpiresAt float64 `json:"exp,omitempty"`
	IssuedAt  float64 `json:"iat,omitempty"`
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Second * 2
