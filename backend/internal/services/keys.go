package services

import (
	"context"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/generators"
	"github.com/crypto-sign/internal/repositories"
	"github.com/sirupsen/logrus"
)

type KeyService struct {
	logger *logrus.Logger
	repo   *repositories.KeyRepository
}

type PublicKeyCreationRequest struct {
	PublicKey string `json:"public_key"`
}

func NewKeyService(logger *logrus.Logger, client clients.Client, keyGenerator *generators.KeysGenerator) *KeyService {
	repo := repositories.NewKeyRepository(logger, client, keyGenerator)
	return &KeyService{
		logger: logger,
		repo:   repo,
	}
}

func (k *KeyService) GetGeneratedKeys(ctx context.Context, userID string) (*domains.Keys, error) {
	keys, err := k.repo.GenerateAndGetInsertedKeys(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("generate and get inserted keys: %w", err)
	}
	return keys, nil
}

func (k *KeyService) GetKeysByUserID(ctx context.Context, userID string) (*domains.Keys, error) {
	keys, err := k.repo.GetKeys(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get keys: %w", err)
	}
	return keys, nil
}

func (k *KeyService) CreatePublicKey(ctx context.Context, request *PublicKeyCreationRequest) error {
	return nil
}
