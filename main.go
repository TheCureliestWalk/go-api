package main

import (
	"fmt"
	"net/http"

	"github.com/TheCureliestWalk/go-api/middleware"
	"github.com/TheCureliestWalk/go-api/routes/users"
	"github.com/TheCureliestWalk/go-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("Server starting...")
	_, sql, err := services.Connect()
	defer sql.Close()
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("Connected to DB.")
	}
	//middlewares
	r.Use(middleware.Cors())

	//activate routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is up and running!",
			"success": true,
		})
	})
	r.GET("/user", users.Read)
	r.GET("/user/:id", users.ReadByID)
	r.POST("/user", users.Create)
	r.DELETE("/user/:id", users.DeleteByID)
	r.PUT("/user/:id", users.UpdateByID)

	//Run the server
	err = r.Run(":5000")
	if err != nil {
		return
	}
}
