package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Flight struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Airline      string             `json:"airline" bson:"airline"`
	FlightNumber string             `json:"flight_number" bson:"flight_number"`
	Origin       string             `json:"origin" bson:"origin"`
	Destiny      string             `json:"destiny" bson:"destiny"`
	PlaneId      string             `json:"plane_id" bson:"plane_id"`
}
