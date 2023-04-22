package handlers

import (
	"fmt"
	"github.com/crypto-sign/internal/services"
	"net/http"
)

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
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs [post]
func (h *DocsHandler) Create(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// GetAvailable docs
// @Summary get available docs for user
// @Description  get available docs for user
// @Tags         docs
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs/available [post]
func (h *DocsHandler) GetAvailable(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
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
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /docs/{doc_id} [get]
func (h *DocsHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}
