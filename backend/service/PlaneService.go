package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var PlaneModel model.Plane

func GetPlane(ctx *gin.Context){
	plane_id := ctx.Param("plane_id")

	res, err := PlaneModel.GetPlaneInfo(plane_id)

	if err == mongo.ErrNoDocuments {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find flight",
		})
		return
	}else if err != nil {
		errorMesage := fmt.Sprintf("Could not find flight due to %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errorMesage,
		})
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Plane found succesfully",
		"plane": res,
	})
}

func SaveNewPlaneToDb(ctx *gin.Context){
	var flight model.Plane

	err := ctx.ShouldBindJSON(&flight)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Could not add new flight",
		})
		return
	}

	newFlight := model.Plane{
		ID: primitive.NewObjectID(),
		PlaneId: flight.PlaneId,
		AirplaneModel: flight.AirplaneModel,
		AvailableSeats: flight.AvailableSeats,
	}
	
	res, err := flight.Save(&newFlight)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not add new flight",
		})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{
		"message":"New flight added", 
		"plane": res,
	})
}
