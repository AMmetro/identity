package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AMmetro/identity/models"
	"github.com/AMmetro/identity/utils"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
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

func formatValidationErrors(err error) []string {
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		out := make([]string, 0, len(vErrs))
		for _, ve := range vErrs {
			field := ve.Field() // short name of the struct field
			switch ve.Tag() {
			case "required":
				out = append(out, fmt.Sprintf("%s is required", field))
			default:
				out = append(out, fmt.Sprintf("%s failed on the '%s' tag", field, ve.Tag()))
			}
		}
		return out
	}
	return []string{err.Error()}
}

func createEvent(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization") // Header = map[string][]string
	// authHeader := context.GetHeader("Authorization") // updated wrapper with Gin
	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}
	token := authHeader
	if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
		token = strings.TrimSpace(authHeader[7:])
	}

	userId, err := utils.Verifytoken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	event.UserID = userId
	// Use ShouldBindJSON for binding validation
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": formatValidationErrors(err)})
		return
	}

	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id", "error": err.Error()})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "not found event with id", "error": err.Error()})
		fmt.Printf("Not found event with id %d: %s\n", eventId, err.Error())
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errors": formatValidationErrors(err)})
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

	err = deletedEvent.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})

}
