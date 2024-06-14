package model

import (
	"context"
	"log"
	"time"

	"github.com/jap102321/flight-system/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Airline      string             `json:"airline" bson:"airline"`
	FlightNumber string             `json:"flight_number" bson:"flight_number"`
	Origin       string             `json:"origin" bson:"origin"`
	Destiny      string             `json:"destiny" bson:"destiny"`
	PlaneId      string             `json:"plane_id" bson:"plane_id"`
	DateOfDeparture time.Time `json:"date_of_departure" bson:"date_of_departure"`
	DateOfArrival time.Time `json:"date_of_arrival" bson:"date_of_arrival"`
	Passengers []primitive.ObjectID `json:"passengers" bson:"passengers"`
	Price int `json:"price" bson:"price"`
}


func (f Flight) GetAllFlights() (interface{}, error){
	ctx := context.Background()
    query, err := db.DB.Collection("flight").Find(ctx, bson.D{{}})

	if err != nil{
		return nil, err
	}

	 
	var flights []bson.M

	if err = query.All(context.TODO(), &flights); err != nil{
		return nil,err
	}

	return flights, nil
}

func (f Flight) GetFlightByFlightNumber(fNumber string) (*primitive.M, error){
	var flight bson.M
	filter := bson.M{"flight_number": fNumber}

	err := db.DB.Collection("flight").FindOne(context.TODO(), filter).Decode(&flight)

	if err != nil {
		return nil, err
	}

	return &flight, nil
}

func (f *Flight) Save(flightToAdd interface{}) (interface{}, error){

	res, err := db.DB.Collection("flight").InsertOne(context.TODO(), flightToAdd)

	if err != nil{
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (f Flight) DeleteEvent(flight_number string) (interface{}, error){
	filter := bson.M{"flight_number": flight_number}


	delRes, err := db.DB.Collection("flight").DeleteOne(context.TODO(), filter)

	if err != nil{
		return nil, err
	}

	return delRes, err
}

func (f Flight) GetFlightByRoute(origin, destiny string) ([]primitive.M, error){
	filter := bson.M{}
	if origin != ""{
		filter["origin"] = origin
	}
	if destiny != ""{
		filter["destiny"] = destiny
	}

	cursor, err := db.DB.Collection("flight").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var flights []bson.M
	
	if err = cursor.All(context.TODO(), &flights); err != nil {
		return nil, err
	}

	return flights, nil
}