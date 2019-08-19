package models

// db2struct --host mysql -d muscle_development -t user_details --package models --struct UserDetail -p --user root --guregu --gorm --json

import (
	null "gopkg.in/guregu/null.v3"
)

// UserDetail UserDetail
type UserDetail struct {
	// ユーザー作成段階では作成されていないテーブルなの全部null許可している
	ID           null.Int    `gorm:"primary_key" json:"id"`
	Height       null.Float  `json:"height"`
	ShortMessage null.String `json:"short_message"`
	UserID       null.Int    `json:"user_id"`
	Weight       null.Float  `json:"weight"`
	CreatedAt    null.Time   `json:"created_at"`
	UpdatedAt    null.Time   `json:"updated_at"`
	DeletedAt    *null.Time  `json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserDetail) TableName() string {
	return "user_details"
}
