package main

import (
	"example.com/m/initializers"
	"example.com/m/models"
)

func init(){
	initializers.LoadEnvVariables()
initializers.ConnectToDB() 
}

func main(){
		 // Migrate the schema
initializers.DB.AutoMigrate(&models.Post{},&models.User{})
	
}