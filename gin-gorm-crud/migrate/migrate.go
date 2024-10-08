package main

import (
	"github.com/VaheMuradyan/Tasks/go-crud/initializers"
	"github.com/VaheMuradyan/Tasks/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
