package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/service"
)

func Routes(server *gin.Engine) {

	flightRoutes := server.Group("/flights")
	flightRoutes.GET("/", service.GetAllFlights)
	flightRoutes.GET("/:flight-number", service.GetFlightByFlightNumber)
	flightRoutes.POST("/", service.SaveFlight)
	flightRoutes.DELETE("/:flight-number", service.DeleteFlight)

	planeRoutes := server.Group("/plane")
	planeRoutes.GET("/:plane_id", service.GetPlane)
	planeRoutes.POST("/", service.SaveNewPlaneToDb)
	planeRoutes.PUT("/:plane_id", service.UpdateAvailSeats)

	customerRoutes := server.Group("/customer")
	customerRoutes.POST("/", service.CreateCustomer)
	customerRoutes.POST("/bulk", service.CreateCustomersBulk)
	customerRoutes.GET("/:document", service.CreateCustomersBulk)

	reservationRoutes := server.Group("/reservation")
	reservationRoutes.POST("/:flight-number", service.SaveReservation)
}