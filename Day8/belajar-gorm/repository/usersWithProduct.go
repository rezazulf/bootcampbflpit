package repository

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"fmt"
)

func GetUsersWithProduct() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}
