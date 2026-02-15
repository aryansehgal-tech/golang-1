package main

import (
	// "net/http"

	"go-tutorial/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H {
	// 		"message": "pong from gin",
	// 	})
	// })

	// router.GET("/me/:id/:newId", func(c *gin.Context) {
	// 	var id = c.Param("id")
	// 	var new_id = c.Param("newId")
	// 	c.JSON(http.StatusOK, gin.H {
	// 		"user_id": id,
	// 		"user_new_id": new_id,
	// 	})
	// })

	// router.POST("/me", func(c *gin.Context){

	// 	type MeRequest struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json: "password"`
	// 	}
	// 	var meRequest MeRequest

	// 	if err := c.BindJSON(&meRequest); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"email": meRequest.Email,
	// 		"password": meRequest.Password,
	// 	})
	// })

	// router.PUT("/me", func(c *gin.Context){

	// 	type MeRequest struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json: "password"`
	// 	}
	// 	var meRequest MeRequest

	// 	if err := c.BindJSON(&meRequest); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})

	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"email": meRequest.Email,
	// 		"password": meRequest.Password,
	// 	})
	// })

	// router.PATCH("/me", func(c *gin.Context){

	// 	type MeRequest struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json: "password"`
	// 	}
	// 	var meRequest MeRequest

	// 	if err := c.BindJSON(&meRequest); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})

	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"email": meRequest.Email,
	// 		"password": meRequest.Password,
	// 	})
	// })

	// router.DELETE("/me/:id", func(c *gin.Context){
		
	// 	var id = c.Param("id")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"id":id,
	// 		"message":"Deleted",
	// 	})
		
	// })

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run()
}
