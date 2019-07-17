package models

// db2struct --host mysql -d muscle_development -t users --package models --struct User -p --user root --guregu --gorm

import (
	null "gopkg.in/guregu/null.v3"
	"time"
)

// User User
type User struct {
	Birthday  null.Time   `gorm:"column:birthday"`
	CreatedAt time.Time   `gorm:"column:created_at"`
	Email     string      `gorm:"column:email"`
	ID        int         `gorm:"column:id;primary_key"`
	Image     string      `gorm:"column:image"`
	Nickname  string      `gorm:"column:nickname"`
	Sex       null.String `gorm:"column:sex"`
	Status    null.String `gorm:"column:status"`
	UID       string      `gorm:"column:uid"`
	UpdatedAt time.Time   `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
