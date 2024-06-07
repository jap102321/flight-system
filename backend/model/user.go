package model

import (
	"context"
	"errors"

	"github.com/jap102321/flight-system/db"
	"github.com/jap102321/flight-system/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	Email        string        `bson:"email" json:"email"`
	Password     string        `bson:"password" json:"password"`
	Reservations []Reservation `bson:"reservations" json:"reservations"`
	IsAdmin bool `bson:"isAdmin" json:"isAdmin"`
}

func getUserByEmail(u *User) User{
	var user User
	filter := bson.M{"email": &u.Email}
	_ = db.DB.Collection("user").FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func (u *User) CreateUser(newUser User) (*mongo.InsertOneResult, error){
	user := getUserByEmail(&newUser)	
	
	if !user.ID.IsZero(){
		return nil, errors.New("email already exists")
	}

	res, err := db.DB.Collection("user").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *User) LogUser() (User, error){
	user := getUserByEmail(u)
	validPassword  := utils.DecryptPassword(u.Password, user.Password)

	if !validPassword{
		return User{}, errors.New("invalid credentials")
	}
	
	return user, nil
}

func (u *User) GetUserById(id primitive.ObjectID) error{
	filter := bson.M{"_id": id}

	err := db.DB.Collection("user").FindOne(context.TODO(), filter).Decode(&u)

	return err
}

func (u *User) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error){
	filter := bson.M{"_id": id}

	res, err := db.DB.Collection("user").DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}