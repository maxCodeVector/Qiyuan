package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

var dbLock sync.Mutex
var dbConnMap map[string]*gorm.DB


func GetConnFromDB(dbPath string) *gorm.DB {
	if dbConnMap == nil{
		dbConnMap = make(map[string]*gorm.DB)
	}
	_, ok := dbConnMap[dbPath ]
	if !ok {
		dbLock.Lock()
		_, ok := dbConnMap[dbPath ]
		if !ok {
			dbConnMap[dbPath] = openDBConn(dbPath)
		}
		dbLock.Unlock()
	}
	dbConn, _ := dbConnMap[dbPath ]
	return dbConn
}

func openDBConn(dbPath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
