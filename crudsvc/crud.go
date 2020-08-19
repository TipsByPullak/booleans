package crudsvc

import (
	"booleans/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//CreateBool : Used to create a new entry in the DB
func CreateBool(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input types.InputBoolean //fetch the values from the user
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newBool := types.Boolean{ID: uuid.New(), Value: input.Value, Key: input.Key}
		if err := db.Create(&newBool).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, newBool)
	}
}

//FetchBool : Used to fetch an entry from the DB
func FetchBool(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id1 := c.Param("id")       //fetch the uuid of the boolean to be updated from the request
		id, err := uuid.Parse(id1) //check if the uuid is valid
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var bool types.Boolean
		if err := db.Where("ID = ?", id).First(&bool).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bool)
	}
}

//UpdateBool : Used to change an entry in the DB
func UpdateBool(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id1 := c.Param("id")       //fetch the uuid of the boolean to be updated from the request
		id, err := uuid.Parse(id1) //check if the uuid is valid
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//fetch the new values from the user
		var input types.InputBoolean
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Check if an entry with the provided UUID even exists in the DB
		var bool types.Boolean
		if err := db.Where("ID = ?", id).First(&bool).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Perform the actual operation to update the record
		data := types.Boolean{ID: id, Value: input.Value, Key: input.Key}
		if err := db.Model(&data).Where("ID = ?", id).Updates(gin.H{"Value": input.Value, "Key": input.Key}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

//DeleteBool : Delete an entry from the DB
func DeleteBool(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id1 := c.Param("id")       //fetch the uuid of the boolean to be updated from the request
		id, err := uuid.Parse(id1) //check if the uuid is valid
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Check if an entry with the provided UUID even exists in the DB
		var bool types.Boolean
		if err := db.Where("ID = ?", id).First(&bool).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Perform the actual deletion
		var data types.Boolean
		data.ID = id
		if err := db.Delete(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, "Success")
	}
}
