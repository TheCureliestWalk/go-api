package users

import (
	"app/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	db, sql, err := database.Connect()
	defer sql.Close()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
	}

	users := []UserModel{}
	db.Table("user").Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"error":   nil,
		"success": true,
	})
}

func ReadByID(c *gin.Context) {
	db, sql, err := database.Connect()
	defer sql.Close()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
	}
	user := []UserModel{} // or -> var user []UserModel
	id := c.Param("id")
	db.Table("user").Where("id = ?", id).Find(&user)
	fmt.Println(user)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"error":   nil,
			"success": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "NOT DATA FOUND!",
		})

	}

}
