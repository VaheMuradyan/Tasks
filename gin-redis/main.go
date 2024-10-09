package main

import (
	"go-redis/controllers"
	"go-redis/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToRedis()
}

func main() {
	r := gin.Default()
	r.POST("/book", controllers.SetAuthorAndBook)
	r.GET("/book", controllers.GetBook)
	r.GET("/books", controllers.GetAllBooks)
	r.DELETE("/book", controllers.DeleteBook)

	r.Run()
}
