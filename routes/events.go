package routes

import (
	"net/http"
	"strconv"

	"example.com/booking/models"
	"github.com/gin-gonic/gin"
)

func getEventById(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with that id"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the data to the server"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {

	idParam := context.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(id)
	if err != err {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with that id"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unnauthorized action"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update, operation failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	idParam := context.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id format"})
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(id)
	if err != err {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with that id"})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unnauthorized action"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete, operation failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted event"})
}
