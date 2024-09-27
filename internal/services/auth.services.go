package services

import (
	"echo-crate/internal/models"
	"echo-crate/internal/repository"
	"net/http"

	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Callback(user *models.User) error
	Logout(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
}

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *AuthService) Callback(user *models.User) error {
	exists, err := s.userRepo.UserExists(user)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	err = s.userRepo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *AuthService) Authenticate(w http.ResponseWriter, r *http.Request) {
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}
