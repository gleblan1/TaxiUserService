package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
	"strconv"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (models.JwtToken, error)
	SignUp(name, phoneNumber, email, password string) (string, error)
	LogOut(ctx context.Context, session, id int) error
	ValidateToken(ctx context.Context, tokenString string) (models.JwtToken, error)
	Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error)
}

func (s *Service) Login(ctx context.Context, phone, password string) (models.JwtToken, error) {
	var tokens models.JwtToken
	passwordFromDB, userId, err := s.repo.GetData(phone)
	if err != nil {
		return models.JwtToken{}, err
	}
	isPasswordCorrect, err := utils.ComparePassword(passwordFromDB, password)
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

func (s *Service) SignUp(name, phoneNumber, email, password string) (models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	existingUserErr := s.repo.CheckUserData(name, phoneNumber, email)
	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot create user: %w", existingUserErr)
	}
	return s.repo.SignUp(name, phoneNumber, email, hashedPassword)
}

func (s *Service) LogOut(ctx context.Context, session, id int) error {
	return s.repo.LogOut(ctx, session, id)
}

func (s *Service) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	claims, err := utils.ExtractClaims(tokenString)
	if err != nil {
		return false, err
	}
	userId := claims.Audience
	sessionId := claims.Session
	accessToken, err := s.repo.ValidateToken(ctx, userId, sessionId)
	if err != nil {
		return false, err
	}
	if accessToken == tokenString {
		return true, nil
	}
	return false, nil
}

func (s *Service) Refresh(ctx context.Context, refreshTokenString string) (models.JwtToken, error) {
	token := models.JwtToken{}
	claims, err := utils.ExtractClaims(refreshTokenString)
	if err != nil {
		return token, err
	}
	userId := claims.Audience
	sessionId := claims.Session
	tokensFromRedis := s.repo.GetRefreshToken(ctx, userId, sessionId)
	if tokensFromRedis == refreshTokenString {

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
