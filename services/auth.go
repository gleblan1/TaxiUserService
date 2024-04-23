package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

type Auth interface {
	Login(ctx context.Context, requestBody config.LoginRequest) (models.JwtToken, error)
	SignUp(ctx context.Context, requestBody config.RegisterRequest) (models.User, error)
	LogOut(ctx context.Context, request config.LogoutRequest) error
	ValidateToken(ctx context.Context, tokenString string) (bool, error)
	Refresh(ctx context.Context, requestBody config.RefreshRequestBody) (models.JwtToken, error)
}

func (s *Service) Login(ctx context.Context, requestBody config.LoginRequest) (models.JwtToken, error) {
	var tokens models.JwtToken
	passwordFromDB, userId, err := s.repo.GetData(requestBody.PhoneNumber)
	if err != nil {
		return models.JwtToken{}, err
	}
	isPasswordCorrect, err := utils.ComparePassword(passwordFromDB, requestBody.Password)
	if err != nil {
		return models.JwtToken{}, errors.New("wrong password")
	}
	session := utils.CreateRandomInt()
	if isPasswordCorrect {
		accessToken, refreshToken, err := utils.GenerateTokens(strconv.Itoa(session), strconv.Itoa(userId))
		if err != nil {
			return models.JwtToken{}, err
		}
		intId, intSession := strconv.Itoa(userId), strconv.Itoa(session)
		err = s.repo.SetTokens(ctx, accessToken, refreshToken, intId, intSession)
		if err != nil {
			return models.JwtToken{}, err
		}
		tokens = models.JwtToken{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		return tokens, nil
	}
	return models.JwtToken{}, nil
}

func (s *Service) SignUp(ctx context.Context, requestBody config.RegisterRequest) (models.User, error) {
	hashedPassword, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		return models.User{}, err
	}
	existingUserErr := s.repo.CheckUserData(requestBody.Name, requestBody.PhoneNumber, requestBody.Email)
	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot create user: %w", existingUserErr)
	}
	return s.repo.SignUp(requestBody.Name, requestBody.PhoneNumber, requestBody.Email, hashedPassword)
}

func (s *Service) LogOut(ctx context.Context, request config.LogoutRequest) error {
	return s.repo.LogOut(ctx, request.SessionId, request.UserId)
}

func (s *Service) Refresh(ctx context.Context, requestBody config.RefreshRequestBody) (models.JwtToken, error) {
	token := models.JwtToken{}
	claims, err := utils.ExtractClaims(requestBody.RefreshToken)
	if err != nil {
		return token, err
	}
	userId := claims.Audience
	sessionId := claims.Session
	tokensFromRedis := s.repo.GetRefreshToken(ctx, userId, sessionId)
	if tokensFromRedis == requestBody.RefreshToken {

		accessToken, refreshToken, err := utils.GenerateTokens(sessionId, userId)

		if err != nil {
			return models.JwtToken{}, err
		}

		token.AccessToken = accessToken
		token.RefreshToken = refreshToken

		err = s.repo.SetTokens(ctx, accessToken, refreshToken, userId, sessionId)
		if err != nil {
			return models.JwtToken{}, err
		}
	} else {
		return models.JwtToken{}, errors.New("refresh token is invalid")
	}
	return token, nil
}
