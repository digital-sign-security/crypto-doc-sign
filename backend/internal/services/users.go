package services

import (
	"github.com/sirupsen/logrus"
)

type UserService struct {
	logger *logrus.Logger
}

func NewUserService(logger *logrus.Logger) *UserService {
	return &UserService{
		logger: logger,
	}
}
