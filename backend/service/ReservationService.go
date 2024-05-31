package service

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func createCustomer (){
	
}


func SaveReservation(ctx *gin.Context) {
	var request struct {
		Reservation model.Reservation `json:"reservation"`
		Customers []model.Customer `json:"customers"`
	}
	
	f_number := ctx.Param("flight-number")
  
	if err := ctx.ShouldBindJSON(&request); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Could not realize the petition.",
		})
		return
	}

	resN, err := request.Reservation.GetReservationData(5)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something went wrong",
		})
		return
	}

	customerIds := make([]primitive.ObjectID, len(request.Customers))


	newReservation := model.Reservation{
		ID: primitive.NewObjectID(),
		ReservationId: strings.ToUpper(resN),
		DateOfDeparture: request.Reservation.DateOfDeparture,
		DateOfArrival: request.Reservation.DateOfArrival,
		FlightNumber: f_number,
	}

	res, err := newReservation.Save(&newReservation)



	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not create reservation",
		})
		return
	}

	err = newReservation.AddReservationToFlight(f_number, newReservation)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not add reservation to flight",
			"error": err.Error(),
		})
	}

	
	for i, customer := range request.Customers{
		customer.ID = primitive.NewObjectID()
		customer.ReservationNumber = newReservation.ReservationId
		customer.Age = time.Now().Year() - customer.DateOfBirth.Year()
		customer.AgeCategory = GetCustomerAgeCategory(customer)
		customerIds[i] = customer.ID
		_, err := customer.SaveClient(customer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":"Could not add client to the reservation",
			})
			return
		}
	}

	filter := bson.M{"_id": newReservation.ID}
	update := bson.M{"$set": bson.M{"customer_ids" : customerIds}}

	err = request.Reservation.UpdateReservation(filter, update)

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Reservation created",
		"res": res.InsertedID,
	})

}