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

var (
	alreadyLogOutErr       = errors.New("already log out")
	invalidRefreshTokenErr = errors.New("refresh token is invalid")
	wrongPasswordErr       = errors.New("wrong password")
)

func (s *Service) SignIn(ctx context.Context, requestBody requests.SignInRequest) (models.JwtToken, error) {
	var tokens models.JwtToken
	user, err := s.repo.GetUserByPhone(ctx, requestBody.PhoneNumber)
	if err != nil {
		return tokens, err
	}
	password := user.Password
	userId := user.Id
	isPasswordCorrect, err := utils.ComparePassword(password, requestBody.Password)
	if err != nil {
		return models.JwtToken{}, wrongPasswordErr
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
	existingUserErr := s.repo.CheckUserData(ctx, requestBody.Name, requestBody.PhoneNumber, requestBody.Email)
	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot create user: %w", existingUserErr)
	}
	signUpData := models.User{
		Name:        requestBody.Name,
		PhoneNumber: requestBody.PhoneNumber,
		Email:       requestBody.Email,
		Password:    hashedPassword,
	}
	userId, err := s.repo.CreateUser(ctx, signUpData)
	if err != nil {
		return models.User{}, err
	}
	user, err := s.repo.GetUserById(ctx, userId)
	return models.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		Rating:      user.Rating,
	}, nil
}

func (s *Service) SignOut(ctx context.Context, request requests.LogoutRequest) error {
	exists, err := s.repo.FindTokens(ctx, strconv.Itoa(request.UserId)+"."+strconv.Itoa(request.SessionId))
	if err != nil {
		return fmt.Errorf("log out: %w", err)
	}
	if exists == 1 {
		s.repo.DeleteTokens(ctx, strconv.Itoa(request.UserId)+"."+strconv.Itoa(request.SessionId))
	} else {
		return fmt.Errorf("log out: %w", alreadyLogOutErr)
	}
	return nil
}

func (s *Service) RefreshTokens(ctx context.Context, requestBody requests.RefreshTokensRequest) (models.JwtToken, error) {
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
		return models.JwtToken{}, invalidRefreshTokenErr
	}
	return token, nil
}
