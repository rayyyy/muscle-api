package models

// db2struct --host mysql -d muscle_development -t user_details --package models --struct UserDetail -p --user root --guregu --gorm --json

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

// UserDetail UserDetail
type UserDetail struct {
	CreatedAt    time.Time   `gorm:"column:created_at" json:"created_at"`
	Height       null.Float  `gorm:"column:height" json:"height"`
	ID           int         `gorm:"column:id;primary_key" json:"id"`
	ShortMessage null.String `gorm:"column:short_message" json:"short_message"`
	UpdatedAt    time.Time   `gorm:"column:updated_at" json:"updated_at"`
	UserID       int         `gorm:"column:user_id" json:"user_id"`
	Weight       null.Float  `gorm:"column:weight" json:"weight"`
}

// TableName sets the insert table name for this struct type
func (u *UserDetail) TableName() string {
	return "user_details"
}
