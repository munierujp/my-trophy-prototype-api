package model

import (
	"my-trophy-prototype-api/types/date"
)

// Trophy is model for `trophies` table
type Trophy struct {
	Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AchievedOn  date.Date `json:"achieved_on"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"-"`
}
