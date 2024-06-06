package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"github.com/jap102321/flight-system/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(ctx *gin.Context){
	var user model.User

	err := ctx.ShouldBindJSON(&user)


	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not recieve data",
		})
		return
	}

	hashedPass, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not has password",
		})
		return
	}

	newUser := model.User{
		ID: primitive.NewObjectID(),
		Email: user.Email,
		Password: hashedPass,
	}


	_, err = user.CreateUser(newUser)

	if err != nil {
		errorMessage := fmt.Sprintf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":errorMessage,
		})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{
		"message":"User created succesfully",
	})
}

func LogIn(ctx *gin.Context){
	var user model.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}
	
	verifiedUser, err := user.LogUser()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJWTToken(user.Email, verifiedUser.ID)

	if err != nil {
		errorMessage := fmt.Sprintf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":errorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login succesfully",
		"token": token,
	})
}