package main

import (
	"belajar-gorm/database"
)

func main() {
	database.StartDB()

	// ======================User========================//
	// repository.CreateUser("rezazulfikarsyah@gmail.com")
	// repository.GetUserById(1)
	// repository.GetAllUsers()
	// repository.UpdateUserById(1, "rezazulfff@gmail.com")

	// ====================Product======================//
	// repository.CreateProduct(1, "YOLO", "YOMZ")
	// repository.DeleteProductById(3)
	// ================Product & Users==================//
	// repository.GetUsersWithProduct()

}
