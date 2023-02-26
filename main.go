package main

import (
	"example.com/m/initializers"
	"example.com/m/routes"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() { 
	// r := gin.Default()
	route:= routes.AddRoute()
	route.Run() // listen and serve on 0.0.0.0:8080
}