package controllers

import (
	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var data[]models.Post
	initializers.DB.Find(&data)
	c.JSON(200, gin.H{
			"data":data ,
		})
	}