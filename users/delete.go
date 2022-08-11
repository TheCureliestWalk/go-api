package users

import (
	"app/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteByID(c *gin.Context) {
	db, sql, err := database.Connect()
	defer sql.Close()
	if err != nil {
		fmt.Printf("Error Connection: %v/n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
	}

	id := c.Param("id")
	var user UserModel
	c.Bind(&user)
	db.Table("user").Where("id = ?", id).Delete(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully!",
		"success": true,
	})

}
