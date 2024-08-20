package main

import (
	"log/slog"
	"net/http"

	"API.com/api-test/db"
	"API.com/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	//Configura o Engine: configura um servidor HTTP
	server := gin.Default()

	// GET quer dois argumentos: 1) O caminho para o qual a solicitação deve ser enviada
	// 2) Handlers: Função que será executada se essa solicitação GET for enviada para /events
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// Um ponteiro para a variável de evento é passado para ShoulBindJSON 
	// internamente o gin olha o body da solicitação de entrada e armazena os dados na variável event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		slog.Default().Info("got nice error", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}