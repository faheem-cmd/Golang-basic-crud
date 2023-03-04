package routes

import (
	"example.com/m/controllers"
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
)



func AddRoute()*gin.Engine {
	route := gin.Default()
	authenticated := route.Group("/")
    authenticated.Use(middleware.RequireAuth())
	route.POST("/posts",controllers.PostsCreate)
	route.GET("/posts",controllers.GetPosts)
	route.GET("/posts/:id",controllers.GetPostById)
	route.PUT("/posts/:id",controllers.UpdatePost)
	route.DELETE("/posts/:id",controllers.DeletePost)
	route.POST("/signup/",controllers.SignUp)
	route.POST("/login/",controllers.Login)
	authenticated.GET("/validate",controllers.Valid)







	return route
}