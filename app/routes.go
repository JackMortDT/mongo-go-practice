package app

import "my-rest-api/controllers"

func routes() {
	router.GET("/person/:id", controllers.GetPerson)
	router.POST("/person", controllers.CreatePerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)
	router.GET("/people", controllers.GetAllPeople)
}
