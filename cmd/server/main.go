package main

import (
	"fmt"

	"github.com/AMmetro/identity/handlers"
	"github.com/AMmetro/identity/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	handlers.RegisterEventsRouts(server)

	server.Run(":8080")

	fmt.Println("Server running on port 8080")
}
