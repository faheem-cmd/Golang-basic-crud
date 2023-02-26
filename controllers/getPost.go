package controllers

import (
	"fmt"

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

func GetPostById(c *gin.Context){
	id := c.Param("id")
		var data models.Post
		initializers.DB.Find(&data,id)
		fmt.Println(data.ID,"hei")
		if data.ID != 0{
        c.JSON(200, gin.H{
			"data":data ,
		})
		}else{
			 c.JSON(200, gin.H{
			"message":"No items" ,
		})
		}
		

	}