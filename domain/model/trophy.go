package model

// Trophy is model for `trophies` table
type Trophy struct {
	Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"-"`
}
