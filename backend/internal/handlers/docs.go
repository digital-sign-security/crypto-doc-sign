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

func (h *DocsHandler) Create(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *DocsHandler) GetAvailable(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *DocsHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}
