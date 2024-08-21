package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) { 
	// GET quer dois argumentos: 1) O caminho para o qual a solicitação deve ser enviada
	// 2) Handlers: Função que será executada se essa solicitação GET for enviada para /events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}