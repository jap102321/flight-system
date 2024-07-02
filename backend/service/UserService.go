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
		IsAdmin: false,
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

	token, err := utils.GenerateJWTToken(user.Email, verifiedUser.ID, user.IsAdmin)

	if err != nil {
		errorMessage := fmt.Sprintf("Error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":errorMessage,
		})
		return
	}

	visibleDataForFE := map[string]interface{}{
		"email": verifiedUser.Email,
		"isAdmin": verifiedUser.IsAdmin,
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login succesfully",
		"token": token,
		"userInfo": visibleDataForFE,
	})
}

func DeleteUser(ctx *gin.Context){
	 var user model.User
	 userId, exists := ctx.Get("userId")
	 paramId := ctx.Param("id")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(paramId)
	
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	if err = user.GetUserById(id); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not find this userId .",
		})
		return
	}

	if userId != user.ID{
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not delete this user due to auth issues.",
		})
		return
	}	

	deleteResult, err := user.DeleteUser(user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not delete this user due to internal error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"User deleted",
		"res": deleteResult,
	})
}