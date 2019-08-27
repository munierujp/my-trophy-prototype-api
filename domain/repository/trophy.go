package repository

import (
	"my-trophy-prototype-api/domain/model"
)

// TrophyRepository is repository interface for Trophy
type TrophyRepository interface {
	Find(*model.Trophy) ([]model.Trophy, error)
}
