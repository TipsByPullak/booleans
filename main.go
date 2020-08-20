package main

import (
	"booleans/crudsvc"
	"booleans/db"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	svPath := "127.0.0.1:3306"
	db := db.StartConn(svPath)
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
