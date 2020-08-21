package main

import (
	"booleans/crudsvc"
	"booleans/db"
	"flag"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbPath := flag.String("db", "127.0.0.1:3306", "Used to provide the path to the MySQL server")
	dbUsr := flag.String("usr", "root", "The username used to access the DB")
	dbPass := flag.String("pass", "", "The passweord to the database")
	flag.Parse()
	db := db.StartConn(*dbPath, *dbUsr, *dbPass)
	defer db.Close()
	router := gin.Default()

	// v1 := router.Group("/booleans")
	// {
	router.POST("/", crudsvc.CreateBool(db))
	router.GET("/:id", crudsvc.FetchBool(db))
	router.DELETE("/:id", crudsvc.DeleteBool(db))
	router.PATCH("/:id", crudsvc.UpdateBool(db))
	router.Run()
}
