package db

import (
	"booleans/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Driver to be used to connect to our DB
)

//StartConn : Function called to initialise a connection to the DB
func StartConn(svPath string, usr string, pass string) *gorm.DB {
	if pass != "" {
		pass = ":" + pass
	}
	dbPath := usr + pass + "@tcp(" + svPath + ")/mysql?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(dbPath)
	db, err := gorm.Open("mysql", dbPath)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&(types.Boolean{}))
	return db
}
