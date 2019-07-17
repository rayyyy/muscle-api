package models

// db2struct --host mysql -d muscle_development -t user_details --package models --struct UserDetail -p --user root --guregu --gorm

import (
	"time"
	null "gopkg.in/guregu/null.v3"
)

// UserDetail UserDetail
type UserDetail struct {
	Bedtime        null.Time   `gorm:"column:bedtime"`
	BloodType      null.String `gorm:"column:blood_type"`
	CreatedAt      time.Time   `gorm:"column:created_at"`
	DrinkingHabits null.String `gorm:"column:drinking_habits"`
	GetUpTime      null.Time   `gorm:"column:get_up_time"`
	Height         null.Float  `gorm:"column:height"`
	ID             int         `gorm:"column:id;primary_key"`
	ShortMessage   null.String `gorm:"column:short_message"`
	SmokingHabit   null.String `gorm:"column:smoking_habit"`
	UpdatedAt      time.Time   `gorm:"column:updated_at"`
	UserID         int         `gorm:"column:user_id"`
	Weight         null.Float  `gorm:"column:weight"`
}

// TableName sets the insert table name for this struct type
func (u *UserDetail) TableName() string {
	return "user_details"
}
