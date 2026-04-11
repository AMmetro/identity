package routes

import (
	"net/http"
	"strconv"

	"github.com/AMmetro/identity/services"
	"github.com/gin-gonic/gin"
)

type UpdateRegistrationStatusInput struct {
	Status string `json:"status" binding:"required"`
}

func getRegistrations(c *gin.Context) {
	registrations, err := services.GetAllRegistrations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"cant`t get registrations ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, registrations)
}

func updateRegistrationStatus(c *gin.Context) {
	userId := c.GetInt64("userId")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid event id"})
		return
	}

	var updatedStatus UpdateRegistrationStatusInput
	if err := c.ShouldBindJSON(&updatedStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err = services.UpdateRegistrationStatus(eventId, userId, updatedStatus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
