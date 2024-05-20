package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/utils"
)

var userIsDeletedErr = errors.New("user is deleted")

func (s *Service) ValidateToken(ctx context.Context, tokenString string) (bool, error) {
	claims, err := utils.ExtractClaims(tokenString)

	if err != nil {
		return false, err
	}

	userId := claims.Audience
	sessionId := claims.Session

	intId, _ := strconv.Atoi(userId)
	if s.repo.IsUserDeleted(ctx, intId) {
		return false, userIsDeletedErr
	}

	accessToken, err := s.repo.ValidateToken(ctx, userId, sessionId)
	if err != nil {
		return false, err
	}

	if accessToken == tokenString {
		return true, nil
	}
	return false, nil
}
