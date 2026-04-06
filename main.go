package main

import (
	"fmt"

	"github.com/AMmetro/identity/db"
	"github.com/AMmetro/identity/routes"
	"github.com/gin-gonic/gin"
	// "github.com/rest-api/identity/db"
	// "github.com/rest-api/identity/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterEventsRouts(server)

	server.Run(":8080")

	fmt.Println("Server running on port 8080")
}
