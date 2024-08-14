package main

import (
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
	server.POST("events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// Um ponteiro para a variável de evento é pasado para ShoulBindJSON 
	// internamente o gin olha o body da solicitação de entrada e armazena os dados na variável event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}