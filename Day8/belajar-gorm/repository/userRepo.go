package repository

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}

func GetUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data : %+v \n", user)
}

func GetAllUsers() {
	db := database.GetDB()

	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("All Users Data:")
	for _, user := range users {
		fmt.Printf("User Data: %+v\n", user)
	}
}

func UpdateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update users' email: %+v", user.Email)
}
