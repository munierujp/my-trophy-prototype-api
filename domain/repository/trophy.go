package repository

import (
	"my-trophy-prototype-api/domain/model"
)

// TrophyRepository is repository interface for Trophy
type TrophyRepository interface {
	Create(*model.Trophy) error
	Find(*model.Trophy) ([]model.Trophy, error)
	FindByID(uint) (*model.Trophy, error)
	Delete(uint) error
}
