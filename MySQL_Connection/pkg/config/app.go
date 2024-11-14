package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	d, err := gorm.Open("mysql", "Database URL")
	if err != nil {
		panic(err)
	}
	db = d
}

func getDb() *gorm.DB {
	return db
}
