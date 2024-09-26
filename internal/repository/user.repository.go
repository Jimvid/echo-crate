package repository

import (
	"echo-crate/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (s *UserRepository) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserRepository) UserExists(user *models.User) (bool, error) {
	var userFound models.User
	result := s.db.First(&userFound, "user_id = ?", user.UserID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}
