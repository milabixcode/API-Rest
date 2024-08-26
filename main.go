package main

import (
	"API.com/db"
	"API.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	//Configura o Engine: configura um servidor HTTP
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8081")
}
