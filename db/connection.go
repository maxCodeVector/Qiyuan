package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

var dbLock sync.Mutex

type DBConn struct {
	name string
	db *gorm.DB
}


var conn *DBConn

func GetDBConn(path string) * DBConn{
	if conn == nil {
		dbLock.Lock()
		if conn == nil {
			conn.db = GetConnFromDB(path)
			conn.name = path
		}
		dbLock.Unlock()
	}
	return conn
}


func GetConnFromDB(dbpath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}


