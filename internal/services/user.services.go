package services

import (
	"echo-crate/internal/models"
	"echo-crate/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *UserService) CreateUser(user *models.User) error {

	err := s.repo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
