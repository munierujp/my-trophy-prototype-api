package repository

import (
	"my-trophy-prototype-api/domain/model"
)

// UserRepository is repository interface for User
type UserRepository interface {
	Find(*model.User) ([]model.User, error)
	FindByID(uint) (*model.User, error)
}
