package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var FlightModel model.Flight


func GetFlightByRoute(ctx *gin.Context){
	origin := ctx.Param("origin")
	destiny := ctx.Param("destiny")

	flights, err := FlightModel.GetFlightByRoute(origin, destiny)
	
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not fetch flights",
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message":"Flights fetched correctly",
		"flights": flights,
	})

}

func GetAllFlights(context *gin.Context) {
	flights, err := FlightModel.GetAllFlights()

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not fetch data of flights",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Flights fetched succesfully",
		"flights" : flights,
	})
}

func GetFlightByFlightNumber(context *gin.Context){
	var flight model.Flight
	flightNumber := context.Param("flight-number")

	flightFetched, err := flight.GetFlightByFlightNumber(strings.ToUpper(flightNumber))	

	if err  != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch flight",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":"Flight fetched succesfully",
		"flight": flightFetched,
	})
}


func SaveFlight(cx *gin.Context){
	var flight model.Flight
	var plane model.Plane

	err := cx.ShouldBindJSON(&flight)

	if err != nil {
		cx.JSON(http.StatusBadRequest,gin.H{
			"message": "Could not parse req data",
		} )
		return
	}

	_, err = plane.GetPlaneInfo(flight.PlaneId)

	if err == mongo.ErrNoDocuments{
		cx.JSON(http.StatusNotFound, gin.H{
			"message":"Plane does not exist",
		})
		return
	}else if err != nil{
		cx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not find error",
		})
		return
	}
	
	if flight.DateOfDeparture.IsZero() || flight.DateOfArrival.IsZero() || flight.Price == 0 {
		cx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid date or price",
		})
		return
	}

	newFlight := model.Flight{
		ID: primitive.NewObjectID(),
		Airline: flight.Airline,
		FlightNumber: strings.ToUpper(flight.FlightNumber),
		DateOfDeparture: flight.DateOfDeparture,
		DateOfArrival: flight.DateOfArrival,
		Origin: flight.Origin,
		Destiny:flight.Destiny,
		PlaneId: flight.PlaneId,
		Passengers: flight.Passengers,
		Price: flight.Price,
	}

	res, err := flight.Save(&newFlight)

	if err != nil {
		cx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save the flight",
		})
		return
	}

	cx.JSON(http.StatusCreated, gin.H{
		"message":"flight created",
		"flight": res,
	})
}


func DeleteFlight(cx *gin.Context){
	id := cx.Param("flight-number")
	res, err := FlightModel.DeleteEvent(id)

	if err == mongo.ErrNoDocuments{
		cx.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find flight",
		})
		return
	}else if err != nil{
		errorMessage := fmt.Sprintf("Could not delete flight due to error %v", err)
		cx.JSON(http.StatusInternalServerError, gin.H{
			"message": errorMessage,
		})
		return
	}

	cx.JSON(http.StatusOK, gin.H{
		"message":"Flight deleted",
		"flightDelete": res,
	})
}
