package middlewares

import (
	"net/http"
	"strings"

	"github.com/AMmetro/identity/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization") // Header = map[string][]string
	// authHeader := context.GetHeader("Authorization") // updated wrapper with Gin
	if authHeader == "" {
		// context.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}
	token := authHeader
	if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
		token = strings.TrimSpace(authHeader[7:])
	}

	userId, err := utils.Verifytoken(token)
	if err != nil {
		// context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
