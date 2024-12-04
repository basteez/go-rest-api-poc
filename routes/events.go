package routes

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"bstz.it/rest-api/models"
	"bstz.it/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		utils.HandleHttpError(err, "Could not fetch events", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusOK, events)
}

func getById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleHttpError(err, "Could not extract id from path", http.StatusBadRequest, context)
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.HandleHttpError(err, "Event not found", http.StatusNotFound, context)
			return
		}
		utils.HandleHttpError(err, "Could not fetch event", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		utils.HandleHttpError(err, "Could not parse request data", http.StatusBadRequest, context)
		return
	}

	err = event.Save()

	if err != nil {
		utils.HandleHttpError(err, "Could not save event", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusCreated, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleHttpError(err, "Could not extract id from path", http.StatusBadRequest, context)
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.HandleHttpError(err, "Event not found", http.StatusNotFound, context)
			return
		}
		utils.HandleHttpError(err, "Could not fetch event", http.StatusInternalServerError, context)
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		utils.HandleHttpError(err, "Could not parse request data", http.StatusBadRequest, context)
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		utils.HandleHttpError(err, "Error updating event", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusNoContent, nil)

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleHttpError(err, "Could not extract id from path", http.StatusBadRequest, context)
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.HandleHttpError(err, "Event not found", http.StatusNotFound, context)
			return
		}
		utils.HandleHttpError(err, "Could not fetch event", http.StatusInternalServerError, context)
		return
	}

	err = event.Delete()
	if err != nil {
		utils.HandleHttpError(err, "Could not delete event", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
