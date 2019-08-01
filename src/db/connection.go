package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func GetConnFromDB() *gorm.DB {
	if db == nil {
		newDb, err := gorm.Open("sqlite3", "test.sqlite")
		if err != nil {
			panic("failed to connect database")
		}
		db = newDb
	}
	return db
}

func CloseDB()  {
	if db != nil {
		db.Close()
	}
}

