package users

import (
	"app/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateByID(c *gin.Context) {
	db, sql, err := database.Connect()
	defer sql.Close()
	if err != nil {
		fmt.Printf("Error Connection: %v/n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
	}

	id := c.Param("id")
	var user UserModel
	c.Bind(&user)
	db.Table("user").Where("id = ?", id).Updates(user)
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Data Updated!",
		"success": true,
	})

}
