package services

import (
	"github.com/sirupsen/logrus"
)

type KeyService struct {
	logger *logrus.Logger
}

func NewKeyService(logger *logrus.Logger) *KeyService {
	return &KeyService{
		logger: logger,
	}
}
