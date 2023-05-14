package services

import (
	"github.com/crypto-sign/internal/domains"
	"github.com/sirupsen/logrus"
)

type CreateDocumentMessageRequest struct {
}

type DocService struct {
	logger *logrus.Logger
}

func NewDocService(logger *logrus.Logger) *DocService {
	return &DocService{
		logger: logger,
	}
}

func (d *DocService) GetAvailableDocumentsForUser() ([]*domains.Doc, error) {
	return nil, nil
}

func (d *DocService) GetDocumentByID(docID string) (*domains.Doc, error) {
	return &domains.Doc{
		ID:              "",
		HashDS:          "",
		Theme:           "",
		DecryptedText:   "",
		SenderUserID:    "",
		RecipientUserID: "",
	}, nil
}

func (d *DocService) CreateDocument(request *CreateDocumentMessageRequest) error {
	return nil
}
