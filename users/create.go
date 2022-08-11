package users

import (
	"app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db, sql, err := database.Connect()
	defer sql.Close()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	var user UserModel // No need to []UserModel
	c.Bind(&user)
	db.Table("user").Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully.",
		"success": true,
	})
}
