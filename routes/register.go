package routes

import (
	"net/http"
	"strconv"

	"example.com/booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	idParam := context.Param("id")
	eventId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to find event with provided id"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to register event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered event!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	idParam := context.Param("id")
	eventId, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to find event with provided id"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration!"})
}
