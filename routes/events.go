package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AMmetro/identity/models"
	"github.com/AMmetro/identity/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id", "error": err.Error()})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	var event models.Event
	event.UserID = userId

	// Use ShouldBindJSON for binding validation
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatValidationErrors(err)})
		return
	}

	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	event.RegisterForEvent(userId)
	context.JSON(http.StatusCreated, gin.H{"message": "event created and registered", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "not found event with id", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you can't update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatValidationErrors(err)})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var newEvent models.Event
	// 	err = context.BindJSON(&newEvent)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// var updatedEvent models.Event
	// existEvent.Name = newEvent.Name
	// existEvent.Description = newEvent.Description
	// existEvent.Location = newEvent.Location
	// existEvent.DateTime = newEvent.DateTime
	// existEvent.UserID = eventId

	// err = existEvent.UpdateEvent()
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }

	context.JSON(http.StatusAccepted, gin.H{"message": "Event updated successfully!", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id", "error": err.Error()})
		return
	}

	deletedEvent, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "not found event with id"})
		fmt.Printf("Not found event with id %d: %s\n", eventId, err.Error())
		return
	}

	userId := context.GetInt64("userId")
	if deletedEvent.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you can't delete this event"})
		return
	}

	err = deletedEvent.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})

}
