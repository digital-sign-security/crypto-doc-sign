package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/crypto-sign/internal/services"
	"net/http"
)

type UsersHandler struct {
	service *services.UserService
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserSignUpResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserSignInResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UsersListResponse struct {
	Items  []*UserResponse `json:"items"`
	Amount int             `json:"amount"`
}

func NewUsersHandler(service *services.UserService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

// GetListOfUsers
// @Summary get all users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
//
//	@Success      200         {object}  UsersListResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users [get]
func (h *UsersHandler) GetListOfUsers(w http.ResponseWriter, r *http.Request) {
	handle := func() (*UsersListResponse, error) {
		users, err := h.service.GetListOfUsers()
		if err != nil {
			return nil, fmt.Errorf("get list of users: %w", err)
		}

		var items []*UserResponse

		for _, item := range users {
			items = append(items, &UserResponse{
				Username: item.Username,
				Email:    item.Email,
			})
		}

		return &UsersListResponse{
			Items:  items,
			Amount: len(items),
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("usersHandler.GetListOfUsers: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}

}

// SignUp
// @Summary user sign up
// @Description  user sign up
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body services.SignUpRequest true "auth params"
//
//	@Success      200         {object}  UserSignUpResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-up [post]
func (h *UsersHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	handle := func() (*UserResponse, error) {
		var requestBody services.SignUpRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			return nil, fmt.Errorf("cannot decode request body: %w", err)
		}

		user, err := h.service.SignUp(requestBody)
		if err != nil {
			return nil, fmt.Errorf("sign up: %w", err)
		}

		return &UserResponse{
			Username: user.Username,
			Email:    user.Email,
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("usersHandler.SignUp: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}

// SignIn
// @Summary user sign in
// @Description  user sign in
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body services.SignInRequest true "auth params"
//
//	@Success      200         {object}  UserSignInResponse
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-in [post]
func (h *UsersHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	handle := func() (*UserResponse, error) {
		var requestBody services.SignInRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			return nil, fmt.Errorf("cannot decode request body: %w", err)
		}

		user, err := h.service.SignIn(requestBody)
		if err != nil {
			return nil, fmt.Errorf("sign in: %w", err)
		}

		return &UserResponse{
			Username: user.Username,
			Email:    user.Email,
		}, nil
	}

	response, err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("usersHandler.SignIn: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return // TODO: FMT
	}
}

// SignOut
// @Summary user sign out
// @Description  user sign out
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body services.SignOutRequest true "auth params"
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-out [post]
func (h *UsersHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	handle := func() error {
		var requestBody services.SignOutRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			return fmt.Errorf("cannot decode request body: %w", err)
		}

		err = h.service.SignOut(requestBody)
		if err != nil {
			return fmt.Errorf("sign out: %w", err)
		}

		return nil
	}

	err := handle()
	if err != nil {
		http.Error(w, fmt.Sprintf("usersHandler.SignOut: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
