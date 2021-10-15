package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/routes"
)

func main() {
	server := gin.Default()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "9000"
	}

	//allowing cors policy
	server.Use(cors.Default())

	//choosing driver database
	connection.Driver("postgres")
	defer connection.GetConnection().Close()
	routes.StartRouting(server)

	//initial migrations
	//connection.StartMigrations()

	server.Run(":" + PORT)
}
