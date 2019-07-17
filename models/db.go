package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetDB DBを取得
func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:pass@tcp(mysql:3306)/muscle_development?charset=utf8&parseTime=True&loc=Local")
	return db, err
}
