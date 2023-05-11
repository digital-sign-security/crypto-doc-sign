package services

import (
	"fmt"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/repositories"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo   *repositories.UserRepository
	logger *logrus.Logger
}

type SignUpRequest struct {
}

type SignInRequest struct {
}

type SignOutRequest struct {
}

func NewUserService(logger *logrus.Logger) *UserService {
	return &UserService{
		logger: logger,
	}
}

func (u *UserService) GetListOfUsers() ([]*domains.User, error) {
	users, err := u.repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	return users, nil
}

func (u *UserService) SignUp(request SignUpRequest) (*domains.User, error) {
	user, err := u.repo.CreateUser()
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	// TODO: create jwt token

	return user, nil
}

func (u *UserService) SignIn(request SignInRequest) (*domains.User, error) {
	user, err := u.repo.GetUser()
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	// TODO: check jwt token

	return user, nil
}

func (u *UserService) SignOut(request SignOutRequest) error {
	_, err := u.repo.PatchUser()
	if err != nil {
		return fmt.Errorf("patch user: %w", err)
	}
	// TODO: remove jwt token

	return nil
}
