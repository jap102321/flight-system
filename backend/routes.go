package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/service"
)

func Routes(server *gin.Engine) {
	server.GET("/flights", service.GetAllFlights)
	server.GET("/flights/:flight-number", service.GetFlightByFlightNumber)
	server.POST("/flights", service.SaveFlight)
	server.DELETE("/flights/:flight-number", service.DeleteFlight)
}