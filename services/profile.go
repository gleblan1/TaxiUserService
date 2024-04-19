package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/GO-Trainee/GlebL-innotaxi-userservice/models"
)

type IProfileService interface {
	GetAccountInfo(id int) (models.UserInfo, error)
	UpdateProfile(id int, newData models.PatchRequest) models.UserInfo
	DeleteProfile(id int) error
}

func (s *Service) GetAccountInfo(id int) (info models.UserInfo, error error) {
	if s.repo.CheckIsUserDeleted(id) {
		return models.UserInfo{}, errors.New("user is deleted")
	}
	return s.repo.GetAccountInfo(id)
}

func (s *Service) UpdateProfile(id int, newData models.PatchRequest) (models.UserInfo, error) {
	if s.repo.CheckIsUserDeleted(id) {
		return models.UserInfo{}, errors.New("user is deleted")
	}
	var newUser models.UserInfo
	oldData, err := s.repo.GetAccountInfo(id)
	if err != nil {
		return newUser, err
	}
	existingUserErr := s.repo.CheckUserData(newData.Name, newData.PhoneNumber, newData.Email)
	if existingUserErr != nil {
		return models.UserInfo{}, fmt.Errorf("cannot update user: %w", existingUserErr)
	}
	if newData.Name != oldData.Name && newData.Name != "" {
		err := s.repo.UpdateUsername(newData.Name, id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	if newData.PhoneNumber != oldData.PhoneNumber && newData.PhoneNumber != "" {
		err := s.repo.UpdatePhoneNumber(newData.PhoneNumber, id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	if newData.Email != oldData.Email && newData.Email != "" {
		err := s.repo.UpdateEmail(newData.Email, id)
		if err != nil {
			return models.UserInfo{}, err
		}
	}
	newUser, err = s.repo.GetAccountInfo(id)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *Service) DeleteProfile(ctx context.Context, id int) error {
	if s.repo.CheckIsUserDeleted(id) {
		return errors.New("user is deleted")
	}
	if err := s.repo.UpdateDeletedStatus(id); err != nil {
		return err
	}
	if err := s.repo.DeleteAllSessions(ctx, id); err != nil {
		return err
	}
	return nil
}
