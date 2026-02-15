package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func main() {
	server := gin.Default()

	// Simple test route
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Get all users
	server.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, users)
	})

	// Create a new user
	server.POST("/user", func(ctx *gin.Context) {
		var newUser User

		// Bind JSON to struct
		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		// Basic validation
		if newUser.Name == "" || newUser.Age <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user data",
			})
			return
		}

		// Add user to slice
		users = append(users, newUser)

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "user created",
			"user":    newUser,
		})
	})

	server.Run(":8080")
}
