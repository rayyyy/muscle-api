package models

import (
	"muscle-api/funcs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetDB DBを取得
func GetDB() (*gorm.DB, error) {
	user := funcs.GetEnv("DATABASE_USER", "root")
	pass := funcs.GetEnv("DATABASE_PASS", "pass")
	host := funcs.GetEnv("DATABASE_HOST", "mysql")
	database := funcs.GetEnv("DATABASE_NAME", "muscle_development?")
	url := user + ":" + pass + "@tcp(" + host + ":3306)/" + database + "charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	return db, err
}
