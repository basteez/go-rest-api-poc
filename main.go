package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"bstz.it/rest-api/configuration"
	"bstz.it/rest-api/db"
	"bstz.it/rest-api/models"
	"bstz.it/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration.LoadConfig()

	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getById)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

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
