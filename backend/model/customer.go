package model

import (
	"context"
	"time"

	"github.com/jap102321/flight-system/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	ID primitive.ObjectID `json:"_id" bson:"_id" `
	Name string `json:"name" bson:"name" binding:"required"`
	LastName string `json:"last_name" bson:"last_name" binding:"required"`
	AgeCategory string `json:"age_category" bson:"age_category"`
	DateOfBirth time.Time `json:"date_of_birth" bson:"date_of_birth"`
	Age int `json:"age" bson:"age"`
	ReservationNumber string `json:"reservation_number" bson:"reservation_number"`
	DocumentNumber string `json:"document_number" bson:"document_number"`
}




func (c Customer) SaveClient(newClient interface{}) (*mongo.InsertOneResult, error){
	res, err := db.DB.Collection("customer").InsertOne(context.TODO(), newClient)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c Customer)GetCustomersByResNumber(res_number string) (*mongo.Cursor, error){
	filter := bson.M{"reservation_number": res_number}

	cursor, err := db.DB.Collection("customer").Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return cursor, nil
}



func (c Customer) GetAllCustomers() (interface{}, error){

	query, err := db.DB.Collection("customer").Find(context.TODO(), bson.D{{}})

	if err != nil {
		return []Customer{}, nil
	}

	var customers []bson.M

	if err = query.All(context.TODO(), &customers); err != nil {
		return nil, err
	}

	return customers, nil
}

func (c Customer) SaveClientBulk(newClients []interface{}) (*mongo.InsertManyResult, error){
	res, err := db.DB.Collection("customer").InsertMany(context.TODO(), newClients)

	if err != nil{
		return nil, err
	}

	return res, nil
}


