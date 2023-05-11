package repositories

import "github.com/crypto-sign/internal/domains"

type UserRepository struct {
}

func (u *UserRepository) GetUsers() ([]*domains.User, error) {
	return nil, nil
}

func (u *UserRepository) GetUser() (*domains.User, error) {
	return nil, nil
}

func (u *UserRepository) CreateUser() (*domains.User, error) {
	return nil, nil
}

func (u *UserRepository) PatchUser() (*domains.User, error) {
	return nil, nil
}
