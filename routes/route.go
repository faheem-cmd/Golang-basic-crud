package routes

import (
	"net/http"
	"strings"

	"example.com/m/controllers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func requireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        header := c.Request.Header.Get("Authorization")
        if header == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(strings.Replace(header, "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {
            // Parse the token using the secret key
            return []byte("FHKAHFKHAKFHKAHFHKAHFAKHF"), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            // Store the user ID in the context for later use
            c.Set("userID", claims["userID"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort()
            return
        }

        c.Next()
    }
}





func AddRoute()*gin.Engine {
	route := gin.Default()
	authenticated := route.Group("/")
    authenticated.Use(requireAuth())
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