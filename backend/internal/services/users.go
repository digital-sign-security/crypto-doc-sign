package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/crypto-sign/internal/auth"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type UserService struct {
	repo     *repositories.UserRepository
	logger   *logrus.Logger
	hashSalt []byte
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
		logger:   logger,
		repo:     repo,
		hashSalt: []byte("dslflsdmfksmdflrwocwo"), // TODO to env
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
	password := u.getUserHashedPassword(request.Password)

	userModel := &domains.User{
		ID:       "",
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}
	user, err := u.repo.CreateUser(ctx, userModel)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(u.hashSalt)
	if err != nil {
		return nil, fmt.Errorf("token signed string: %w", err)
	}
	u.logger.Infof(tokenString)
	user, err = u.repo.CreateToken(ctx, tokenString, user)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}

	return user, nil
}

func (u *UserService) SignIn(ctx context.Context, request SignInRequest) (*domains.User, error) {
	password := u.getUserHashedPassword(request.Password)
	user, err := u.repo.GetUser(ctx, request.Username, password)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(u.hashSalt)
	if err != nil {
		return nil, fmt.Errorf("token signed string: %w", err)
	}
	u.logger.Infof(tokenString)

	user, err = u.repo.CreateToken(ctx, tokenString, user)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}

	return user, nil
}

func (u *UserService) SignOut(ctx context.Context, request SignOutRequest) error {
	_, err := u.repo.PatchUser(ctx)
	if err != nil {
		return fmt.Errorf("patch user: %w", err)
	}
	// TODO: remove auth token

	return nil
}

func (u *UserService) createJWTToken() {

}

func (u *UserService) getUserHashedPassword(password string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write(u.hashSalt)
	return fmt.Sprintf("%x", pwd.Sum(nil))
}
