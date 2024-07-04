package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/middleware"
	"github.com/jap102321/flight-system/service"
)

func Routes(server *gin.Engine) {


	flightRoutes := server.Group("/flights")
	flightRoutes.GET("/all-flights", service.GetAllFlights)
	flightRoutes.GET("/:flight-number", service.GetFlightByFlightNumber)
	flightRoutes.GET("/route/:origin/:destiny", service.GetFlightByRoute)

	customerRoutes := server.Group("/customer")
	customerRoutes.POST("/", service.CreateCustomer)
	customerRoutes.POST("/bulk", service.CreateCustomersBulk)

	reservationRoutes := server.Group("/reservation")
	reservationRoutes.POST("/:flight-number", service.SaveReservation)
	reservationRoutes.GET("/:flight_number", service.GetReservations)

	userRoutes := server.Group("/user")
	userRoutes.POST("/register", service.CreateUser)
	userRoutes.POST("/login", service.LogIn)

	
	authenticated := server.Group("/auth")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/flight", service.SaveFlight)
	authenticated.GET("/plane/:plane_id", service.GetPlane)
	authenticated.POST("/plane", service.SaveNewPlaneToDb)
	authenticated.DELETE("/user/:id", service.DeleteUser)
	authenticated.DELETE("/customer/:id", service.DeleteCustomers)
	authenticated.DELETE("/:flight-number", service.DeleteFlight)

}