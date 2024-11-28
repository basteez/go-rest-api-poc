package main

import (
	"net/http"

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

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		utils.HandleHttpError(err, "Could not parse request data", http.StatusBadRequest, context)
		return
	}

	event.ID = 1     // TODO remove dummy
	event.UserID = 1 // TODO remove dummy

	err = event.Save()

	if err != nil {
		utils.HandleHttpError(err, "Could not save event", http.StatusInternalServerError, context)
		return
	}

	context.JSON(http.StatusCreated, event)
}
