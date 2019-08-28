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

// Create creates new trophy
func (repo *TrophyRepository) Create(trophy *model.Trophy) error {
	if err := repo.db.Create(&trophy).Error; err != nil {
		return err
	}
	return nil
}

// Find returns trophies match to query
func (repo *TrophyRepository) Find(query *model.Trophy) ([]model.Trophy, error) {
	trophies := []model.Trophy{}
	if err := repo.db.Where(query).Find(&trophies).Error; err != nil {
		return nil, err
	}
	return trophies, nil
}

// FindByID returns trophy matches to id
func (repo *TrophyRepository) FindByID(id uint) (*model.Trophy, error) {
	var trophy model.Trophy
	if err := repo.db.First(&trophy, id).Error; err != nil {
		return nil, err
	}
	return &trophy, nil
}

// Delete deletes user matches to id
func (repo *TrophyRepository) Delete(id uint) error {
	trophy := model.Trophy{}
	trophy.ID = id
	if err := repo.db.Delete(&trophy).Error; err != nil {
		return err
	}
	return nil
}
