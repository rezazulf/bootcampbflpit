package main

import (
	"assignment-2_reza/config"
	"assignment-2_reza/controllers"
	"assignment-2_reza/routers"
)

var PORT = ":3000"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.StartServer(controller).Run(PORT)
}
