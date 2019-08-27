package database

import (
	"my-trophy-prototype-api/domain/model"
	"my-trophy-prototype-api/domain/repository"

	"github.com/jinzhu/gorm"
)

// TrophyRepository is repository for Trophy
type TrophyRepository struct {
	db *gorm.DB
}

// NewTrophyRepository is constructor for TrophyRepository
func NewTrophyRepository(db *gorm.DB) repository.TrophyRepository {
	return &TrophyRepository{db}
}

// Find returns trophies match to query
func (repo *TrophyRepository) Find(query *model.Trophy) ([]model.Trophy, error) {
	trophies := []model.Trophy{}
	if err := repo.db.Where(query).Find(&trophies).Error; err != nil {
		return nil, err
	}
	return trophies, nil
}