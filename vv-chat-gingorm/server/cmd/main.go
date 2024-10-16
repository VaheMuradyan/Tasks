package main

import (
	"log"
	"vvchat/server/db"
	"vvchat/server/internal/user"
	"vvchat/server/internal/ws"
	"vvchat/server/middleware"
	"vvchat/server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("!!! could not initialize database connection: %s", err)
	}

	db := dbConn.GetDB()

	db.AutoMigrate(&user.User{})

	userHandler := user.NewHandler(db)
	middleware1 := middleware.NewMiddleware(db)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	router.InitRouter(userHandler, middleware1, wsHandler)
	router.Start("0.0.0.0:8080")
}
