package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func getDotEnv() (string, error){
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("could not fetch .env token")
	}

	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		return "", errors.New("could not load env")
	}

	return secretKey, nil
}


func GenerateJWTToken(email string, userId primitive.ObjectID) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"uId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	secretKey, _ := getDotEnv()
	

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (primitive.ObjectID, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		secretKey, err := getDotEnv()
		if err != nil {
			return nil, errors.New("could not get secret key")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("token could not be parsed: %v", err)
	}

	if !parsedToken.Valid {
		return primitive.NilObjectID, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.NilObjectID, errors.New("invalid token claims")
	}

	uIdStr, ok := claims["uId"].(string)
	if !ok {
		return primitive.NilObjectID, errors.New("uId claim is missing or not a string")
	}

	uId, err := primitive.ObjectIDFromHex(uIdStr)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid uId: %v", err)
	}

	return uId, nil
}