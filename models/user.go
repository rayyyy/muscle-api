package models

// db2struct --host mysql -d muscle_development -t users --package models --struct User -p --user root --guregu --gorm

type User struct {
	Birthday  time.Time `gorm:"column:birthday"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Email     string    `gorm:"column:email"`
	ID        int       `gorm:"column:id;primary_key"`
	Image     string    `gorm:"column:image"`
	Nickname  string    `gorm:"column:nickname"`
	Sex       string    `gorm:"column:sex"`
	Status    string    `gorm:"column:status"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

