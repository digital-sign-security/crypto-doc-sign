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

// GetGeneratedKeys
// @Summary get generated public and private keys
// @Description  get generated public and private keys for user from server
// @Tags         keys
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys [get]
func (h *KeysHandler) GetGeneratedKeys(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// Post keys
// @Summary post your public key in the system
// @Description  post your public key in the system
// @Tags         keys
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public [post]
func (h *KeysHandler) Post(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// Get keys
// @Summary get your public key
// @Description  get your public
// @Tags         keys
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public [get]
func (h *KeysHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// GetAnotherUserPublicKey
// @Summary get another user public key
// @Description  get another user public key
// @Tags         keys
// @Accept       json
// @Produce      json
// @Param		 user_id path string true "The unique user_id of a User"
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public/{user_id} [get]
func (h *KeysHandler) GetAnotherUserPublicKey(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// GetServerPublicKey
// @Summary get server public key
// @Description  get server public key
// @Tags         keys
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public/server [get]
func (h *KeysHandler) GetServerPublicKey(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}
