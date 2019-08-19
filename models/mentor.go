package models

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

// Mentor Mentor
type Mentor struct {
	ID            int         `gorm:"primary_key" json:"id"`
	UserID        int         `json:"user_id"`
	Title         string      `json:"title"`
	AppealMessage string      `json:"appeal_message"`
	Image1        null.String `json:"image1"`
	Image2        null.String `json:"image2"`
	Image3        null.String `json:"image3"`
	Image4        null.String `json:"image4"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     null.Time   `json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (m *Mentor) TableName() string {
	return "mentors"
}
