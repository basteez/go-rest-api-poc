package main

import (
	"bstz.it/rest-api/configuration"
	"bstz.it/rest-api/db"
	"bstz.it/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration.LoadConfig()

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
