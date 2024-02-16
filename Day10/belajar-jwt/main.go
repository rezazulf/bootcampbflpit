package main

import (
	"belajar-jwt/database"
	"belajar-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
