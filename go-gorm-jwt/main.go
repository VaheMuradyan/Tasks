package main

import (
	"github.com/VaheMuradyan/Tasks/go-gorm-jwt/controllers"
	"github.com/VaheMuradyan/Tasks/go-gorm-jwt/initilazers"
	"github.com/VaheMuradyan/Tasks/go-gorm-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initilazers.LoadEnvVariables()
	initilazers.ConnectToDb()
	initilazers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
