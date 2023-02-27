package routes

import (
	"example.com/m/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoute()*gin.Engine {
	route := gin.Default()
	route.POST("/posts",controllers.PostsCreate)
	route.GET("/posts",controllers.GetPosts)
	route.GET("/posts/:id",controllers.GetPostById)
	route.PUT("/posts/:id",controllers.UpdatePost)
	route.DELETE("/posts/:id",controllers.DeletePost)
	return route
}