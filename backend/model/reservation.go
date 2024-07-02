package model

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/jap102321/flight-system/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reservation struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	ReservationId string `json:"reservation_id" bson:"reservation_id"`
	Customers []primitive.ObjectID `json:"customers_ids" bson:"customers_ids"`
	DateOfDeparture time.Time `json:"date_of_departure" bson:"date_of_departure"`
	DateOfArrival time.Time `json:"date_of_arrival" bson:"date_of_arrival"`
	FlightNumber string `json:"flight_number" bson:"flight_number"`
	Email string `json:"email" bson:"email"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (r Reservation) GetReservationData(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func (r Reservation)GetAllReservations() (*mongo.Cursor, error){
	res, err := db.DB.Collection("reservation").Find(context.TODO(), bson.D{{}})
	
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r Reservation) GetAllReservationsForFlight(flightNumber string) ([]bson.M, error) {
    filter := bson.M{"flight_number": flightNumber}

    pipeline := mongo.Pipeline{
        {{Key: "$match", Value: filter}},
        {{Key: "$lookup", Value: bson.M{
            "from":         "customer",
            "localField":   "customers_ids",
            "foreignField": "_id",
            "as":           "customers",
        }}},
        {{Key: "$project", Value: bson.M{
            "customers_ids": 0,
        }}},
    }

    cursor, err := db.DB.Collection("reservation").Aggregate(context.TODO(), pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var reservations []bson.M
    if err = cursor.All(context.TODO(), &reservations); err != nil {
        return nil, err
    }

    return reservations, nil
}

func (r Reservation) Save(newReservation interface{})(*mongo.InsertOneResult, error){

	res, err := db.DB.Collection("reservation").InsertOne(context.TODO(), newReservation)
	
	if err != nil{
		return nil, err
	}

	return res, nil
}

func (r Reservation) DeleteReservation(res_number string) (*mongo.DeleteResult, error) {
	filter := bson.M{"reservation_id": res_number}
	res, err := db.DB.Collection("reservation").DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r Reservation) AddReservationToFlight(flight_number string, reservation primitive.ObjectID) error{
	filter := bson.M{"flight_number" : flight_number}
	update := bson.M{"$push": bson.M{"passengers" : reservation}}

	fmt.Println(reservation)

	_, err := db.DB.Collection("flight").UpdateOne(context.TODO(), filter, update)

	return err
}

func (r Reservation) UpdateReservation(filter, update primitive.M) error{
	_, err := db.DB.Collection("reservation").UpdateOne(context.TODO(), filter, update)

	return err
}
