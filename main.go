package main

import (
	"fmt"
	"net/http"

	"github.com/AMmetro/identity/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

	fmt.Println("Server running on port 8080")
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = 1
	event.ID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})
}
