package models

// db2struct --host mysql -d muscle_development -t users --package models --struct User -p --user root --guregu --gorm --json

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

// User User
type User struct {
	ID        int         `gorm:"primary_key" json:"id"`
	Birthday  null.Time   `json:"birthday"`
	Email     string      `json:"email"`
	Image     string      `json:"image"`
	Nickname  string      `json:"nickname"`
	Sex       null.String `gorm:"default:'unspecified'" json:"sex"`
	Status    null.String `gorm:"default:'active'" json:"status"`
	UID       string      `json:"uid"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`

	UserDetail UserDetail `json:"user_detail"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
