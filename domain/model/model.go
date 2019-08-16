package model

import (
	"time"
)

// Model is base model definition, which is compatible with "github.com/jinzhu/gorm".Model
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
