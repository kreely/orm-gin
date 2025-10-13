package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kreely/orm-gin/initialisers"
	"github.com/kreely/orm-gin/models"


)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	// Create a post
	// Return it.
	post := models.Post{Title: "Hello", Body: "Post body"}
	result := initialisers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	
		c.JSON(200, gin.H{
			"post": post,
		})
	}
