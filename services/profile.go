package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/requests"
)

func (s *Service) GetAccountInfo(ctx context.Context, req requests.GetAccountInfoRequest) (info models.User, error error) {
	if s.repo.IsUserDeleted(ctx, req.Id) {
		return models.User{}, errors.New("user is deleted")
	}
	return s.repo.GetAccountInfo(req.Id)
}

func (s *Service) UpdateProfile(ctx context.Context, req requests.UpdateProfileRequest) (models.User, error) {
	if s.repo.IsUserDeleted(ctx, req.Id) {
		return models.User{}, errors.New("user is deleted")
	}

	var newUser models.User

	oldData, err := s.repo.GetAccountInfo(req.Id)

	if err != nil {
		return newUser, err
	}
	existingUserErr := s.repo.CheckUserData(ctx, req.Name, req.PhoneNumber, req.Email)

	if existingUserErr != nil {
		return models.User{}, fmt.Errorf("cannot update user: %w", existingUserErr)
	}
	if req.Name != oldData.Name && req.Name != "" {
		err := s.repo.UpdateUsername(req.Name, req.Id)
		if err != nil {
			return models.User{}, err
		}
	}
	if req.PhoneNumber != oldData.PhoneNumber && req.PhoneNumber != "" {
		err := s.repo.UpdatePhoneNumber(req.PhoneNumber, req.Id)
		if err != nil {
			return models.User{}, err
		}
	}
	if req.Email != oldData.Email && req.Email != "" {
		err := s.repo.UpdateEmail(req.Email, req.Id)
		if err != nil {
			return models.User{}, err
		}
	}
	newUser, err = s.repo.GetAccountInfo(req.Id)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *Service) DeleteProfile(ctx context.Context, req requests.DeleteProfileRequest) error {
	if s.repo.IsUserDeleted(ctx, req.Id) {
		return errors.New("user is deleted")
	}
	if err := s.repo.UpdateDeletedStatus(req.Id); err != nil {
		return err
	}
	if err := s.repo.DeleteAllSessions(ctx, req.Id); err != nil {
		return err
	}
	return nil
}
