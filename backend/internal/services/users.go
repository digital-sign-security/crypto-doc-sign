package services

import (
	"context"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/repositories"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo   *repositories.UserRepository
	logger *logrus.Logger
}

type SignUpRequest struct {
	Username string
	Email    string
	Password string
}

type SignInRequest struct {
	Username string
	Password string
}

type SignOutRequest struct {
	Username string
	JWTToken string
}

func NewUserService(logger *logrus.Logger, client clients.Client) *UserService {
	repo := repositories.NewUserRepository(logger, client)
	return &UserService{
		logger: logger,
		repo:   repo,
	}
}

func (u *UserService) GetListOfUsers(ctx context.Context) ([]*domains.UserWithKey, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	return users, nil
}

func (u *UserService) SignUp(ctx context.Context, request SignUpRequest) (*domains.User, error) {
	userModel := &domains.User{
		ID:       "",
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := u.repo.CreateUser(ctx, userModel)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	// TODO: create jwt token

	return user, nil
}

func (u *UserService) SignIn(ctx context.Context, request SignInRequest) (*domains.User, error) {
	user, err := u.repo.GetUser(ctx, request.Username, request.Password)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	// TODO: check jwt token

	return user, nil
}

func (u *UserService) SignOut(ctx context.Context, request SignOutRequest) error {
	_, err := u.repo.PatchUser(ctx)
	if err != nil {
		return fmt.Errorf("patch user: %w", err)
	}
	// TODO: remove jwt token

	return nil
}
