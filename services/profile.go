package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice/config"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type Profile interface {
	GetAccountInfo(ctx context.Context, req config.GetAccountInfoRequest) (models.UserInfo, error)
	UpdateProfile(ctx context.Context, req config.UpdateProfileRequest) (models.UserInfo, error)
	DeleteProfile(ctx context.Context, req config.DeleteProfileRequest) error
}

func (s *Service) GetAccountInfo(ctx context.Context, req config.GetAccountInfoRequest) (info models.UserInfo, error error) {
	if s.repo.CheckIsUserDeleted(req.Id) {
		return models.UserInfo{}, errors.New("user is deleted")
	}
	return s.repo.GetAccountInfo(req.Id)
}

func (s *Service) UpdateProfile(ctx context.Context, req config.UpdateProfileRequest) (models.UserInfo, error) {
	if s.repo.CheckIsUserDeleted(req.Id) {
		return models.UserInfo{}, errors.New("user is deleted")
	}

	var newUser models.UserInfo

	oldData, err := s.repo.GetAccountInfo(req.Id)

	if err != nil {
		return newUser, err
	}
	existingUserErr := s.repo.CheckUserData(req.NewData.Name, req.NewData.PhoneNumber, req.NewData.Email)

	if existingUserErr != nil {
		return models.UserInfo{}, fmt.Errorf("cannot update user: %w", existingUserErr)
	}
	if req.NewData.Name != oldData.Name && req.NewData.Name != "" {
		err := s.repo.UpdateUsername(req.NewData.Name, req.Id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	if req.NewData.PhoneNumber != oldData.PhoneNumber && req.NewData.PhoneNumber != "" {
		err := s.repo.UpdatePhoneNumber(req.NewData.PhoneNumber, req.Id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	if req.NewData.Email != oldData.Email && req.NewData.Email != "" {
		err := s.repo.UpdateEmail(req.NewData.Email, req.Id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	newUser, err = s.repo.GetAccountInfo(req.Id)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *Service) DeleteProfile(ctx context.Context, req config.DeleteProfileRequest) error {
	if s.repo.CheckIsUserDeleted(req.Id) {
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
