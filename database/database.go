package database

import "gorm.io/gorm"

func Initialise() *gorm.DB {
	var db *gorm.DB
	var err error

	if err != nil {
		panic(err)
	}

	return db
}
