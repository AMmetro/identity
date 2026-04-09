package routes

import (
	"net/http"

	"github.com/AMmetro/identity/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// result, err := models.FindByEmail(user.Email)
	isValid, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Not found user with email": err.Error()})
		return
	}

	if isValid {
		context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully!"})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}
