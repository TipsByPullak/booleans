package main

import (
	"booleans/crudsvc"
	"booleans/db"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var svPath string
	fmt.Printf("Please input the IP and the port to the DB server (usage: <IP:PORT>)- ")
	fmt.Scanf("%s", &svPath)
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
