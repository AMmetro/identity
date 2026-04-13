package handlers

import (
	"github.com/AMmetro/identity/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRouts(server *gin.Engine) {

	authenticated := server.Group("/")

	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PATCH("/events/:id/registration", updateRegistrationStatus)

	server.GET("/events", getEvents)
	server.GET("/registration", getRegistrations)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
