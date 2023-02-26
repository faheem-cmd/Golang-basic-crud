package controllers

import (
	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)
 
func PostsCreate(c *gin.Context) {
	// Get data from req body
	var body  struct{
		Body string
		Title string	
	}
	c.Bind(&body)
	if  body.Title == "" ||body.Body == "" {
		 c.JSON(400,gin.H{"status":"failed",})
		 return
	}
	// Create post
	post:=models.Post{Title:body.Title,Body:body.Body }
	result:=initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
		c.JSON(200, gin.H{
			"data":post,
		})
	}

	 

	