package model

import (
	"context"
	"errors"

	"github.com/jap102321/flight-system/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Plane struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	PlaneId string	`json:"plane_id" bson:"plane_id"`
	AirplaneModel string `json:"airplane_model" bson:"airplane_model"`
	AvailableSeats int64 `json:"available_seats" bson:"available_seats"`
}


func (p Plane) GetPlaneInfo(plane_id string) (*primitive.M, error){
	var plane bson.M
	
	err:= db.DB.Collection("plane").FindOne(context.TODO(), bson.D{primitive.E{Key: "plane_id", Value: plane_id}}).Decode(&plane)

	if err != nil {
		return nil, err
	}

	return &plane, err
}


func (p Plane) Save(planeToAdd interface{}) (interface{}, error){
	res, err := db.DB.Collection("plane").InsertOne(context.TODO(), planeToAdd)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p Plane) DeletePlane(flight_number string) (*mongo.DeleteResult, error) {
	filter := bson.M{"flight_number" : flight_number}

	res, err := db.DB.Collection("plane").DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return nil, err 
	}

	return res, nil
}

func (p *Plane)UpdatePlaneSeats(plane_id string, numberOfCustomers int) (*mongo.UpdateResult, error){
	var plane Plane
	res, err := plane.GetPlaneInfo(plane_id)

	if err != nil {
		return nil, err
	}

	id, ok := (*res)["_id"].(primitive.ObjectID)
	if !ok{
		return nil, errors.New("could not find id")
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "available_seats", Value: - numberOfCustomers}} }}


	resUpdate, err := db.DB.Collection("plane").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err 
	}

	return resUpdate, err 
}
