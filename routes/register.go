package routes

import (
	"net/http"
	"strconv"

	"API.com/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event 
	event.ID = eventID

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
