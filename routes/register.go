package routes

import (
	"net/http"
	"strconv"

	"github.com/AMmetro/identity/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	registrationId, err := event.RegisterForEvent(userId)
	if err != nil {
		// check for UNIQUE constraint UNIQUE(event_id, user_id)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Successfully registered",
		"registrationId": registrationId,
	})
}
