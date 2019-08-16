package handler

import "github.com/jinzhu/gorm"

// Handler is handler for request
type Handler struct {
	DB *gorm.DB
}
