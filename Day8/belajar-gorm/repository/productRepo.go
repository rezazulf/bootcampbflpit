package repository

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"fmt"
)

func CreateProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("Now Product Data:", Product)
}

func DeleteProductById(id uint) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}
	fmt.Printf("Product with id %d has been successfully deleted", id)
}
