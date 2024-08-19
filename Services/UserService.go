package Services

import (
	"Golang/Models"
	"Golang/Repository"
	"Golang/utils"
	"errors"
)

type UserService interface {
	AuthenticateUser(login, password string) (*Models.User, string, error)
	RegisterUser(username, login, password string) (*Models.User, string, error)
	ConfirmUserEmail(token string) error
}

type UserServiceImpl struct {
	repository          Repository.UserRepository
	confirmationService EmailConfirmationService
}

func (s UserServiceImpl) AuthenticateUser(login, password string) (*Models.User, string, error) {
	user, err := s.repository.FindByEmail(login)
	if err != nil {
		return nil, "", err
	}

	if !user.EmailConfirmed {
		return nil, "", errors.New("Email is not confirmed")
	}

	if !user.CheckPassword(password) {
		return nil, "", errors.New("Incorrect password")
	}
	token, err := utils.GenerateAuthToken(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s UserServiceImpl) RegisterUser(username, login, password string) (*Models.User, string, error) {
	user, err := s.repository.Create(username, login, password)
	if err != nil {
		return nil, "", err
	}
	token, err := utils.GenerateAuthToken(user)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s UserServiceImpl) ConfirmUserEmail(token string) error {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	user, err := s.repository.FindByEmail(claims.Email)
	if err != nil {
		return err
	}
	user.EmailConfirmed = true
	s.repository.Update(user)

	return nil
}

func NewUserService(
	repository Repository.UserRepository,
	confirmationService EmailConfirmationService,
) UserService {
	return &UserServiceImpl{
		repository:          repository,
		confirmationService: confirmationService,
	}
}
