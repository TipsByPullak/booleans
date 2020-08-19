package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db1, err1 := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
	if err1 != nil {
		panic(err)
	}
	defer db1.Close()

	_, err = db1.Exec("CREATE DATABASE IF NOT EXISTS " + "demo")
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&boolean{})
}

func main() {

	router := gin.Default()

	v1 := router.Group("/booleans")
	{
		v1.POST("/", createBool)
		v1.GET("/:id", fetchBool)
		v1.DELETE("/:id", deleteBool)
		v1.PATCH("/:id", updateBool)
	}
	router.Run()

}

type (
	boolean struct {
		ID    uuid.UUID `json:"id"`
		Value bool      `json:"value"`
		Key   string    `json:"key"`
	}

	inputBoolean struct {
		Value bool   `json:"value"`
		Key   string `json:"key"`
	}
)

func createBool(c *gin.Context) {
	var input inputBoolean //fetch the values from the user
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(input)
		return
	}

	newBool := boolean{ID: uuid.New(), Value: input.Value, Key: input.Key}
	if err := db.Create(&newBool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newBool)
}

func fetchBool(c *gin.Context) {
	id1 := c.Param("id")
	id, err := uuid.Parse(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var bool boolean
	if err := db.Where("ID = ?", id).First(&bool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bool)
	c.JSON(http.StatusOK, bool)
}

func updateBool(c *gin.Context) {
	id1 := c.Param("id") //fetch the uuid of the boolean to be updated
	id, err := uuid.Parse(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fetch the new values from the user
	var input inputBoolean
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := boolean{ID: id, Value: input.Value, Key: input.Key}
	if err := db.Model(&data).Where("ID = ?", id).Updates(gin.H{"Value": input.Value, "Key": input.Key}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func deleteBool(c *gin.Context) {
	id1 := c.Param("id")
	id, err := uuid.Parse(id1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var data boolean
	data.ID = id
	if err := db.Delete(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "Success")
}
