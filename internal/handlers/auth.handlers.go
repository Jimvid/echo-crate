package handlers

import (
	"echo-crate/internal/models"
	"echo-crate/internal/services"
	"echo-crate/internal/sessions"
	pages "echo-crate/internal/views/pages"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(r.Context(), w)
}

func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)

	newUser := &models.User{
		Email:  user.Email,
		UserID: user.UserID,
	}

	if err != nil {
		fmt.Fprintln(w, err)
	}

	err = h.service.Callback(newUser)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.SaveUserToSession(w, r, user)

	if err != nil {
		http.Error(w, "Failed to save user session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	h.service.Logout(w, r)
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	h.service.Authenticate(w, r)
}
