package repositories

import (
	"context"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/generators"
	"github.com/sirupsen/logrus"
)

type KeyRepository struct {
	client clients.Client
	logger *logrus.Logger

	keyGenerator *generators.KeysGenerator
}

func NewKeyRepository(logger *logrus.Logger, client clients.Client, keyGenerator *generators.KeysGenerator) *KeyRepository {
	return &KeyRepository{
		logger:       logger,
		client:       client,
		keyGenerator: keyGenerator,
	}
}

func (k *KeyRepository) GenerateAndGetInsertedKeys(ctx context.Context, userID string) (*domains.Keys, error) {
	keys, err := k.keyGenerator.GenerateKeys(ctx, k.logger)
	if err != nil {
		return nil, fmt.Errorf("generate keys: %w", err)
	}

	q := `
		INSERT INTO public.keys 
		    (public_key, private_key, is_alive, user_id) 
		VALUES 
			($1, $2, $3, $4) 
		RETURNING id
	`
	k.logger.Infof("SQL Query: %s", formatQuery(q))

	err = k.client.QueryRow(ctx, q, keys.PublicKey, keys.PrivateKey, true, userID).Scan(&keys.ID)
	if err != nil {
		k.logger.Errorf("SQL Query: %s - err: %s", formatQuery(q), err)
		return &domains.Keys{}, fmt.Errorf("cannot get query: %w", err)
	}

	return keys, nil
}

func (k *KeyRepository) GetKeys(ctx context.Context, userID string) (*domains.Keys, error) {
	q := `
		SELECT
			id, public_key, private_key, is_alive, user_id
		FROM public.keys
		WHERE
			user_id=$1
	`
	k.logger.Infof("SQL Query: %s", formatQuery(q))

	var keys domains.Keys
	err := k.client.QueryRow(ctx, q, userID).Scan(&keys.ID, &keys.PublicKey, &keys.PrivateKey, &keys.IsAlive, &keys.UserID)
	if err != nil {
		k.logger.Errorf("SQL Query: %s - err: %s", formatQuery(q), err)
		return &domains.Keys{}, fmt.Errorf("cannot get query: %w", err)
	}

	return &keys, nil
}
