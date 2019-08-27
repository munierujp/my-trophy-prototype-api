package database

import (
	"my-trophy-prototype-api/domain/model"
	"my-trophy-prototype-api/domain/repository"

	"github.com/jinzhu/gorm"
)

// UserRepository is repository for User
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is constructor for UserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db}
}

// FindByID returns user matches to id
func (repo *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
