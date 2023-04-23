package handlers

import (
	"fmt"
	"github.com/crypto-sign/internal/services"
	"net/http"
)

type UsersHandler struct {
	service *services.UserService
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
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users [get]
func (h *UsersHandler) GetListOfUsers(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// SignUp
// @Summary user sign up
// @Description  user sign up
// @Tags         users
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-up [post]
func (h *UsersHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// SignIn
// @Summary user sign in
// @Description  user sign in
// @Tags         users
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-in [post]
func (h *UsersHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}

// SignOut
// @Summary user sign out
// @Description  user sign out
// @Tags         users
// @Accept       json
// @Produce      json
//
//	@Success      200         {string}  string "OK"
//	@Failure      400         {string}  string  "Bad Request"
//	@Failure      500         {string}  string  "Internal Server Error"
//
// @Router       /users/sign-out [post]
func (h *UsersHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("%v", r.URL)))
	if err != nil {
		return
	}
}
