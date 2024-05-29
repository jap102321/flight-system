package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var customer model.Customer



func getCustomerAgeCategory(customer model.Customer) string{
	if customer.Age >= 16 {
		customer.AgeCategory = "Adult"
		return customer.AgeCategory
	}else{
		customer.AgeCategory = "Minor"
		return customer.AgeCategory
	}
}

func GetCustomers(ctx *gin.Context){

	res, err := customer.GetAllCustomers()

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch customers",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":"Customers fetched",
		"customers": res,
	})
}

func CreateCustomer(ctx *gin.Context){
	var customer model.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create customer",
		})
		return
	}

	newCustomer := model.Customer{
		ID: primitive.NewObjectID(),
		Name: customer.Name,
		LastName: customer.LastName,
		Age : time.Now().Year() - customer.DateOfBirth.Year(),
		DateOfBirth: customer.DateOfBirth,
		ReservationNumber: customer.ReservationNumber,
	}
	newCustomer.AgeCategory = getCustomerAgeCategory(newCustomer)

	res, err := customer.SaveClient(newCustomer)

	if err!= nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save customer",
		})
		return
	}


	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Created!",
		"customerId": res.InsertedID,
	})
}

func CreateCustomersBulk(ctx *gin.Context){
	var customers []model.Customer

	if err := ctx.ShouldBindJSON(&customers); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not save customers",
		})
		return
	}


	var documents []interface{}

	for _, customer := range customers{
		customer.ID = primitive.NewObjectID()
		customer.Age = time.Now().Year() - customer.DateOfBirth.Year()
		customer.AgeCategory = getCustomerAgeCategory(customer)

		documents = append(documents, customer)
	}

	res, err := customer.SaveClientBulk(documents)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error adding files",
		})
		return
	}
	

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"Saved",
		"customersIDs": res.InsertedIDs,
	})
}