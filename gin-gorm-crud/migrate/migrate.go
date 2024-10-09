package main

import (
	"github.com/VaheMuradyan/Tasks/gin-gorm-crud/initializers"
	"github.com/VaheMuradyan/Tasks/gin-gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
