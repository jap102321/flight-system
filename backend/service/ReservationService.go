package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type requestType struct {
	Reservation model.Reservation `json:"reservation"`
	Customers   []model.Customer  `json:"customers"`
}


func GetReservations(ctx *gin.Context){
	var reservation model.Reservation
	flight_number := ctx.Param("flight_number")

	res, err := reservation.GetAllReservationsForFlight(flight_number)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find reservations",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"reservations": res,
	})

}

func SaveReservation(ctx *gin.Context) {
	var (
		plane   model.Plane
		flight  model.Flight
		request requestType
	)

	fNumber := ctx.Param("flight-number")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not process the request."})
		return
	}

	reservationID, err := request.Reservation.GetReservationData(5)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	newReservation := createNewReservation(request, reservationID, fNumber)
	res, err := newReservation.Save(&newReservation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create reservation"})
		return
	}

	if err := newReservation.AddReservationToFlight(fNumber, newReservation.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not add reservation to flight", "error": err.Error()})
		return
	}

	customerIds := make([]primitive.ObjectID, len(request.Customers))
	if err := addCustomersToReservation(newReservation, customerIds, request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update reservation to add customers"})
		return
	}

	if _, err := updatePlaneSeats(flight, plane, fNumber, len(customerIds)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update plane seats"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Reservation created", "res": res.InsertedID})
}

func createNewReservation(request requestType, reservationID, flightNumber string) model.Reservation {
	return model.Reservation{
		ID:             primitive.NewObjectID(),
		ReservationId:  strings.ToUpper(reservationID),
		DateOfDeparture: request.Reservation.DateOfDeparture,
		DateOfArrival:   request.Reservation.DateOfArrival,
		FlightNumber:    flightNumber,
	}
}

func addCustomersToReservation(newReservation model.Reservation, customerIds []primitive.ObjectID, request requestType) error {
	for i, customer := range request.Customers {
		customer.ID = primitive.NewObjectID()
		customer.ReservationNumber = newReservation.ReservationId
		customer.Age = time.Now().Year() - customer.DateOfBirth.Year()
		customer.AgeCategory = GetCustomerAgeCategory(customer)
		customerIds[i] = customer.ID
		if _, err := customer.SaveClient(customer); err != nil {
			return err
		}
	}

	filter := bson.M{"_id": newReservation.ID}
	update := bson.M{"$set": bson.M{"customers_ids": customerIds}}

	return request.Reservation.UpdateReservation(filter, update)
}

func updatePlaneSeats(flight model.Flight, plane model.Plane, flightNumber string, customersCount int) (*mongo.UpdateResult, error) {
	flightData, err := flight.GetFlightByFlightNumber(flightNumber)
	if err != nil {
		return nil, err
	}

	planeID, ok := (*flightData)["plane_id"].(string)
	if !ok {
		return nil, errors.New("could not get plane id")
	}

	return plane.UpdatePlaneSeats(planeID, customersCount)
}
