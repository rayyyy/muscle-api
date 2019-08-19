package models

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

// MentorPlan MentorPlan
type MentorPlan struct {
	ID        int       `gorm:"primary_key" json:"id"`
	MentorID  int       `json:"mentor_id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	IsValid   int       `json:"is_valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (m *MentorPlan) TableName() string {
	return "mentor_plans"
}
