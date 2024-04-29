package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

func (s *Service) Login(ctx context.Context, requestBody requests.LoginRequest) (models.JwtToken, error) {
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

func (s *Service) SignUp(ctx context.Context, requestBody requests.RegisterRequest) (models.User, error) {
	hashedPassword, err := utils.HashPassword(requestBody.Password)
	if err != nil {
		return models.User{}, err
	}
	existingUserErr := s.repo.CheckUserData(requestBody.Name, requestBody.PhoneNumber, requestBody.Email)
	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot create user: %w", existingUserErr)
	}
	userId, err := s.repo.CreateUser(requestBody.Name, requestBody.PhoneNumber, requestBody.Email, hashedPassword)
	if err != nil {
		return models.User{}, err
	}
	user, err := s.repo.GetUser(userId)
	return models.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Rating:      user.Rating,
	}, nil
}

func (s *Service) LogOut(ctx context.Context, request requests.LogoutRequest) error {
	return s.repo.LogOut(ctx, request.SessionId, request.UserId)
}

func (s *Service) Refresh(ctx context.Context, requestBody requests.RefreshRequestBody) (models.JwtToken, error) {
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
