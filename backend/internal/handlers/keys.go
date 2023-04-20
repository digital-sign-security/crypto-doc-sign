package handlers

import (
	"fmt"
	"github.com/crypto-sign/internal/services"
	"net/http"
)

type KeysHandler struct {
	service *services.KeyService
}

func NewKeysHandler(service *services.KeyService) *KeysHandler {
	return &KeysHandler{
		service: service,
	}
}

func (h *KeysHandler) Post(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *KeysHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *KeysHandler) GetAnotherUserPublicKey(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *KeysHandler) GetServerPublicKey(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

func (h *KeysHandler) GetKeys(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}
