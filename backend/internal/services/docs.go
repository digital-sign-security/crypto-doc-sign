package services

import (
	"github.com/sirupsen/logrus"
)

type DocService struct {
	logger *logrus.Logger
}

func NewDocService(logger *logrus.Logger) *DocService {
	return &DocService{
		logger: logger,
	}
}
