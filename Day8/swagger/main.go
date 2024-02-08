package main

import (
	"swagger/database"
	_ "swagger/docs"
	"swagger/routers"
)

// @Summary Get a car by ID
// @Description Get a car from the database by ID
// @ID get-car-by-id
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} CarResponse
// @Failure 404 {object} ErrorResponse
// @Router /cars/{id} [get]

func main() {
	database.StartDB()
	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}
