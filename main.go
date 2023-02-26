package main

import (
	"example.com/m/controllers"
	"example.com/m/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() { 
	r := gin.Default()
	r.POST("/posts",controllers.PostsCreate)
	r.GET("/posts",controllers.GetPosts)

	r.Run() // listen and serve on 0.0.0.0:8080
}