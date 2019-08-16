package repository

import (
	"my-trophy-prototype-api/domain/model"
)

// UserRepository is repository interface for User
type UserRepository interface {
	FindByID(uint) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
