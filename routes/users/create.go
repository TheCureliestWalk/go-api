package users

import (
	"net/http"

	"github.com/TheCureliestWalk/go-api/models"
	"github.com/TheCureliestWalk/go-api/services"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	db, sql, err := services.Connect()
	defer sql.Close()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	var user models.UserModel // No need to []UserModel
	c.Bind(&user)
	db.Table("user").Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully.",
		"success": true,
	})
}
