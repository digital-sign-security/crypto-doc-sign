package repositories

import (
	"context"
	"fmt"
	"github.com/crypto-sign/internal/clients"
	"github.com/crypto-sign/internal/domains"
	"github.com/sirupsen/logrus"
)

type DocRepository struct {
	client clients.Client
	logger *logrus.Logger
}

func NewDocRepository(logger *logrus.Logger, client clients.Client) *DocRepository {
	return &DocRepository{
		logger: logger,
		client: client,
	}
}

func (d *DocRepository) GetDocument(ctx context.Context, docID string) (*domains.Doc, error) {
	q := `
		SELECT
			id, theme, sender_user_id, recipient_user_id, hash_ds, decrypted_text
		FROM public.docs
		WHERE
			id = $1
	`
	d.logger.Infof("SQL Query: %s", formatQuery(q))

	var doc domains.Doc
	err := d.client.QueryRow(ctx, q, docID).Scan(
		&doc.ID, &doc.Theme, &doc.SenderUserID, &doc.RecipientUserID, &doc.HashDS, &doc.DecryptedText,
	)
	if err != nil {
		d.logger.Errorf("SQL Query: %s - err: %s", formatQuery(q), err)
		return &domains.Doc{}, fmt.Errorf("cannot get query: %w", err)
	}

	return &doc, nil
}

func (d *DocRepository) GetDocumentsList(ctx context.Context, userID string) ([]*domains.Doc, error) {
	q := `
		SELECT 
			id, theme, sender_user_id, recipient_user_id
		FROM public.docs
		WHERE
			recipient_user_id = $1
	`
	d.logger.Infof("SQL Query: %s", formatQuery(q))

	rows, err := d.client.Query(ctx, q, userID)
	if err != nil {
		return nil, err // TODO REHANDLE ERRORS
	}

	docs := make([]*domains.Doc, 0)

	for rows.Next() {
		var doc domains.Doc

		err = rows.Scan(&doc.ID, &doc.Theme, &doc.SenderUserID, &doc.RecipientUserID)
		if err != nil {
			return nil, err
		}

		docs = append(docs, &doc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return docs, nil
}

func (d *DocRepository) CreateDocument(ctx context.Context, doc *domains.Doc) (*domains.Doc, error) {
	q := `
		INSERT INTO public.docs 
		    (hash_ds, decrypted_text, theme, sender_user_id, recipient_user_id) 
		VALUES 
			($1, $2, $3, $4, $5) 
		RETURNING id
	`
	d.logger.Infof("SQL Query: %s", formatQuery(q))

	err := d.client.QueryRow(
		ctx, q, &doc.HashDS, &doc.Theme, &doc.DecryptedText, &doc.SenderUserID, &doc.RecipientUserID,
	).Scan(&doc.ID)
	if err != nil {
		d.logger.Errorf("SQL Query: %s - err: %s", formatQuery(q), err)
		return &domains.Doc{}, fmt.Errorf("cannot get query: %w", err)
	}

	return doc, nil
}
