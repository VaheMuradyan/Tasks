package main

import (
	"github.com/VaheMuradyan/Tasks/go-crud/controllers"
	"github.com/VaheMuradyan/Tasks/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts/:id", controllers.GetPostById)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
