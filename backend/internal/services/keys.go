package services

import (
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/generators"
	"github.com/sirupsen/logrus"
)

type KeyService struct {
	logger       *logrus.Logger
	keyGenerator *generators.KeysGenerator
}

type PublicKeyCreationRequest struct {
	PublicKey string `json:"public_key"`
}

func NewKeyService(logger *logrus.Logger, client clients.Client, keyGenerator *generators.KeysGenerator) *KeyService {
	return &KeyService{
		logger:       logger,
		keyGenerator: keyGenerator,
	}
}

func (k *KeyService) GetGeneratedKeys() (*domains.Keys, error) {
	public, private := k.keyGenerator.GenerateKeys()
	return &domains.Keys{
		ID:         "",
		PublicKey:  string(public),
		PrivateKey: string(private),
		IsAlive:    true,
		UserID:     "",
	}, nil
}

func (k *KeyService) GetKeysByUserID(userID string) (*domains.Keys, error) {
	return &domains.Keys{
		ID:         "",
		PublicKey:  "public" + userID,
		PrivateKey: "private",
		IsAlive:    true,
		UserID:     "",
	}, nil
}

func (k *KeyService) CreatePublicKey(request *PublicKeyCreationRequest) error {
	return nil
}
