package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/middleware"
	"github.com/jap102321/flight-system/service"
)

func Routes(server *gin.Engine) {

	planeRoutes := server.Group("/plane")
	planeRoutes.GET("/:plane_id", service.GetPlane)
	planeRoutes.POST("/", service.SaveNewPlaneToDb)


	flightRoutes := server.Group("/flights")
	flightRoutes.GET("/", service.GetAllFlights)
	flightRoutes.GET("/:flight-number", service.GetFlightByFlightNumber)
	flightRoutes.POST("/", service.SaveFlight)
	flightRoutes.DELETE("/:flight-number", service.DeleteFlight)

	customerRoutes := server.Group("/customer")
	customerRoutes.POST("/", service.CreateCustomer)
	customerRoutes.POST("/bulk", service.CreateCustomersBulk)
	customerRoutes.GET("/:document", service.CreateCustomersBulk)

	reservationRoutes := server.Group("/reservation")
	reservationRoutes.POST("/:flight-number", service.SaveReservation)
	reservationRoutes.GET("/:flight_number", service.GetReservations)

	userRoutes := server.Group("/user")
	userRoutes.POST("/register", service.CreateUser)
	userRoutes.POST("/login", service.LogIn)

	
	authenticated := server.Group("/auth")
	authenticated.Use(middleware.Authenticate)
	authenticated.GET("/flight", service.GetAllFlights)
	

}