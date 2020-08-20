//Package db : Only for implementations in Docker where connecting to the localhost is needed
package db

import (
	"booleans/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Driver to be used to connect to our DB
)

//StartConn : Function called to initialise a connection to the DB
func StartConn(svPath string) *gorm.DB {
	usr := "tipsbypullak"
	pass := "MySQL@1026"
	if pass != "" {
		pass = ":" + pass
	}
	//WARNING: The following line is only for connecting the docker container to the localhost, and doesn't work on Linux for now.
	dbPath := usr + pass + "@tcp(host.docker.internal:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(dbPath)
	db, err := gorm.Open("mysql", dbPath)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&(types.Boolean{}))
	return db
}
