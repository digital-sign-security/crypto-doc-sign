package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/crypto-sign/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type KeysResponse struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type PublicKeyResponse struct {
	PublicKey string `json:"public_key"`
}

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
//	@Success      200         {object}  KeysResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys [get]
func (h *KeysHandler) GetGeneratedKeys(w http.ResponseWriter, r *http.Request) {
	handle := func() (*KeysResponse, error) {
		// TODO get user from jwt_token
		keys, err := h.service.GetGeneratedKeys(r.Context(), "38946d88-3aae-4928-bc62-984ec7543dbb")
		if err != nil {
			return nil, fmt.Errorf("get generated keys: %w", err)
		}

		return &KeysResponse{
			PrivateKey: keys.PrivateKey,
			PublicKey:  keys.PublicKey,
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("KeysHandler.GetGeneratedKeys: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}

// UploadAnotherPublicKey
// @Summary post your public key in the system
// @Description  post your public key in the system
// @Tags         keys
// @Accept       json
// @Produce      json
// @Param request body services.PublicKeyCreationRequest true "public key from user"
//
//	@Success      201         {string}  string  "Created"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public [post]
func (h *KeysHandler) UploadAnotherPublicKey(w http.ResponseWriter, r *http.Request) {
	handle := func() error {
		var p services.PublicKeyCreationRequest

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			return fmt.Errorf("cannot decode request body: %w", err)
		}

		// TODO get user from jwt_token
		err = h.service.CreatePublicKey(r.Context(), &p)
		if err != nil {
			return fmt.Errorf("create public key: %w", err)
		}

		return nil
	}

	err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("KeysHandler.UploadAnotherPublicKey: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUserPublicKey
// @Summary get user public key by user_id
// @Description  get user public key by user_id
// @Tags         keys
// @Accept       json
// @Produce      json
// @Param		 user_id path string true "The unique user_id of a User"
//
//	@Success      200         {object}  PublicKeyResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /keys/public/{user_id} [get]
func (h *KeysHandler) GetUserPublicKey(w http.ResponseWriter, r *http.Request) {
	handle := func() (*PublicKeyResponse, error) {
		userID := chi.URLParam(r, "user_id")
		keys, err := h.service.GetKeysByUserID(r.Context(), userID)
		if err != nil {
			return nil, fmt.Errorf("get keys by user id: %w", err)
		}

		return &PublicKeyResponse{
			PublicKey: keys.PublicKey,
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("KeysHandler.GetUserPublicKey: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}
