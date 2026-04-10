package routes

import (
	"github.com/AMmetro/identity/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventsRouts(server *gin.Engine) {

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
