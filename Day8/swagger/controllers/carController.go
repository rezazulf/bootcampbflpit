package controllers

import (
	"fmt"
	"net/http"
	"swagger/database"
	"swagger/models"

	"github.com/gin-gonic/gin"
)

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /orders [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error Getting Car Datas ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetOneCar godoc
// @Summary Get details for a given Id
// @Description Get details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetOneCars(c *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car
	err := db.First(&carOne, "id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data One ": carOne})
}

// CreateCars godoc
// @Summary Post details for a given Id
// @Description Post details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/ [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()

	var input models.Car

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	carinput := models.Car{
		Merk:     input.Merk,
		TypeCars: input.TypeCars,
		Harga:    input.Harga,
	}
	db.Create(&carinput)
	c.JSON(http.StatusOK, gin.H{"data": carinput})
}

// UpdateCars godoc
// @Summary Update car indetified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars/{id} [PUT]
func UpdateCars(c *gin.Context) {
	carID := c.Param("id")
	var db = database.GetDB()

	var carUpdate models.Car
	err := db.First(&carUpdate, "Id = ?", carID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan Model sebelum melakukan perubahan nilai kolom
	db.Model(&carUpdate).Updates(models.Car{Merk: input.Merk, TypeCars: input.TypeCars, Harga: input.Harga})

	c.JSON(http.StatusOK, gin.H{"data": carUpdate})

}

// DeleteCars godoc
// @Summary Delete car identified by the given Id
// @Description Delete the order corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{Id} [delete]
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()

	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}
	db.Delete(&carDelete)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
