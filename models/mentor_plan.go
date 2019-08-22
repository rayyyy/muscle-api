package models

import (
	null "gopkg.in/guregu/null.v3"
)

// MentorPlan MentorPlan
type MentorPlan struct {
	ID        int         `gorm:"primary_key" json:"id"`
	MentorID  int         `json:"mentor_id"`
	Title     null.String `json:"title"`
	Detail    null.String `json:"detail"`
	Type      null.String `gorm:"default:'once'" json:"type"`
	Price     null.Int    `json:"price"`
	IsValid   null.Int    `gorm:"default:1" json:"is_valid"`
	CreatedAt null.Time   `json:"created_at"`
	UpdatedAt null.Time   `json:"updated_at"`
	DeletedAt null.Time   `json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (m *MentorPlan) TableName() string {
	return "mentor_plans"
}
