package controllers

import (
	"go-tutorial/services"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesServices services.NotesServices
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes()) 
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesServices.GetNotesServices(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesServices.CreateNotesServices(),
		})
	}
}
