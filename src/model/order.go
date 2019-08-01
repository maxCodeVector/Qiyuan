package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Order struct {
	gorm.Model
	OrderId  string
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
}

// set User's table name to be `profiles`
func (Order) TableName() string {
	return "demo_order"
}

func Main() {
	db, err := gorm.Open("sqlite3", "test.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Order{})
	db.Create(&Order{OrderId:"1223234", UserName:"hya", Amount:100, Status:"OK", FileUrl:"www.baidu.com"})
	print("success open db")
}