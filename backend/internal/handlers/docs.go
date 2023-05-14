package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/crypto-sign/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DocumentItemResponse struct {
	Theme           string `json:"theme"`
	SenderUserID    string `json:"sender_id"`
	RecipientUserID string `json:"recipient_id"`
}

type AvailableDocumentsResponse struct {
	Items  []*DocumentItemResponse `json:"items"`
	Amount int                     `json:"amount"`
}

type DocumentResponse struct {
	HashDS          string `json:"hash"`
	Theme           string `json:"Theme"`
	DecryptedText   string `json:"decrypted_text"`
	SenderUserID    string `json:"sender_id"`
	RecipientUserID string `json:"recipient_id"`
}

type DocsHandler struct {
	service *services.DocService
}

func NewDocsHandler(service *services.DocService) *DocsHandler {
	return &DocsHandler{
		service: service,
	}
}

// Create docs
// @Summary create decrypted doc with signature
// @Description  create decrypted doc with signature
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param request body services.CreateDocumentMessageRequest true "document message creation"
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs [post]
func (h *DocsHandler) Create(w http.ResponseWriter, r *http.Request) {
	handle := func() error {
		var requestBody services.CreateDocumentMessageRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			return fmt.Errorf("cannot decode request body: %w", err)
		}

		// TODO get user from jwt_token
		err = h.service.CreateDocument(&requestBody)
		if err != nil {
			return fmt.Errorf("create document: %w", err)
		}

		return nil
	}

	err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("DocsHandler.Create: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAvailable docs
// @Summary get available docs for user
// @Description  get available docs for user
// @Tags         docs
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  AvailableDocumentsResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs/available [get]
func (h *DocsHandler) GetAvailable(w http.ResponseWriter, r *http.Request) {
	handle := func() (*AvailableDocumentsResponse, error) {
		// TODO get user from jwt_token
		documents, err := h.service.GetAvailableDocumentsForUser()
		if err != nil {
			return nil, fmt.Errorf("get available documents for user: %w", err)
		}

		var items []*DocumentItemResponse

		for _, item := range documents {
			items = append(items, &DocumentItemResponse{
				Theme:           item.Theme,
				SenderUserID:    item.SenderUserID,
				RecipientUserID: item.RecipientUserID,
			})
		}

		return &AvailableDocumentsResponse{
			Items:  items,
			Amount: len(items),
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("DocsHandler.GetAvailable: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}

// Get doc
// @Summary get document by doc_id
// @Description  get document by doc_id
// @Tags         docs
// @Accept       json
// @Produce      json
// @Param		 doc_id path string true "The doc_id of a document"
//
//	@Success      200         {object}  DocumentResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs/{doc_id} [get]
func (h *DocsHandler) Get(w http.ResponseWriter, r *http.Request) {
	handle := func() (*DocumentResponse, error) {
		documentID := chi.URLParam(r, "document_id")
		// TODO get user from jwt_token
		document, err := h.service.GetDocumentByID(documentID)
		if err != nil {
			return nil, fmt.Errorf("get document by id: %w", err)
		}
		_ = document

		return &DocumentResponse{
			HashDS:          document.HashDS,
			DecryptedText:   document.DecryptedText,
			Theme:           document.Theme,
			SenderUserID:    document.SenderUserID,
			RecipientUserID: document.RecipientUserID,
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("DocsHandler.Get: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}
