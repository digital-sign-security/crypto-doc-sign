package services

import (
	"context"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/crypto-sign/internal/repositories"
	"github.com/sirupsen/logrus"
)

type CreateDocumentMessageRequest struct {
	HashDS          string `json:"hash"`
	Theme           string `json:"Theme"`
	DecryptedText   string `json:"decrypted_text"`
	SenderUserID    string `json:"sender_id"`
	RecipientUserID string `json:"recipient_id"`
}

type DocService struct {
	logger *logrus.Logger
	repo   *repositories.DocRepository
}

func NewDocService(logger *logrus.Logger, client clients.Client) *DocService {
	repo := repositories.NewDocRepository(logger, client)
	return &DocService{
		logger: logger,
		repo:   repo,
	}
}

func (d *DocService) GetAvailableDocumentsForUser(ctx context.Context, userID string) ([]*domains.Doc, error) {
	documents, err := d.repo.GetDocumentsList(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get documents list: %w", err)
	}
	return documents, nil
}

func (d *DocService) GetDocumentByID(ctx context.Context, docID string) (*domains.Doc, error) {
	document, err := d.repo.GetDocument(ctx, docID)
	if err != nil {
		return nil, fmt.Errorf("get document: %w", err)
	}
	return document, nil
}

func (d *DocService) CreateDocument(ctx context.Context, request *CreateDocumentMessageRequest) error {
	doc := &domains.Doc{
		ID:              "",
		HashDS:          request.HashDS,
		Theme:           request.Theme,
		DecryptedText:   request.DecryptedText,
		SenderUserID:    request.SenderUserID,
		RecipientUserID: request.RecipientUserID,
	}
	_, err := d.repo.CreateDocument(ctx, doc)
	if err != nil {
		return fmt.Errorf("create document: %w", err)
	}
	return nil
}
