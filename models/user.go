package models

// db2struct --host mysql -d muscle_development -t users --package models --struct User -p --user root --guregu --gorm --json

import (
	"time"
	null "gopkg.in/guregu/null.v3"
)

// User User
type User struct {
	ID        int         `gorm:"column:id;primary_key" json:"id"`
	Birthday  null.Time   `gorm:"column:birthday" json:"birthday"`
	Email     string      `gorm:"column:email" json:"email"`
	Image     string      `gorm:"column:image" json:"image"`
	Nickname  string      `gorm:"column:nickname" json:"nickname"`
	Sex       null.String `gorm:"column:sex" json:"sex"`
	Status    null.String `gorm:"column:status" json:"status"`
	UID       string      `gorm:"column:uid" json:"uid"`
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time  `gorm:"column:deleted_at" json:"deleted_at"`

	UserDetail UserDetail `json:"user_detail"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
