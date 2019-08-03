package models

// db2struct --host mysql -d muscle_development -t user_details --package models --struct UserDetail -p --user root --guregu --gorm --json

import (
	"time"

	null "gopkg.in/guregu/null.v3"
)

// UserDetail UserDetail
type UserDetail struct {
	ID           int         `gorm:"column:id;primary_key" json:"id"`
	Height       null.Float  `gorm:"column:height" json:"height"`
	ShortMessage null.String `gorm:"column:short_message" json:"short_message"`
	UserID       int         `gorm:"column:user_id" json:"user_id"`
	Weight       null.Float  `gorm:"column:weight" json:"weight"`
	CreatedAt    time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *null.Time  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserDetail) TableName() string {
	return "user_details"
}
