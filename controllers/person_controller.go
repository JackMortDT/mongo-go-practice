package controllers

import (
	"my-rest-api/domain/person"
	service "my-rest-api/services/person"
	"my-rest-api/utils/error_utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPerson(c *gin.Context) {
	var filter bson.M = bson.M{}

	if c.Param("id") != "" {
		id := c.Param("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	person, err := service.PersonService.GetPerson(filter)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	c.JSON(http.StatusOK, person)
}

func GetAllPeople(c *gin.Context) {
	people, err := service.PersonService.GetAllPeople()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, people)
}

func CreatePerson(c *gin.Context) {
	var person person.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	per, err := service.PersonService.CreatePerson(&person)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, per)
}

func UpdatePerson(c *gin.Context) {
	var person person.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	update := bson.M{
		"$set": person,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	personUpdated, err := service.PersonService.UpdatePerson(bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, personUpdated)
}

func DeletePerson(c *gin.Context) {
	objID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	res, err := service.PersonService.DeletePerson(bson.M{"_id": objID})
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, res)
}
