package services

import (
	"context"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

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
