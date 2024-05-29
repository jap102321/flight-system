package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID primitive.ObjectID
	ReservationId string
	Customers []Customer
}