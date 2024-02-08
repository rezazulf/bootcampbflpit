package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars/:carId", controllers.GetCar)
	router.GET("/cars", controllers.GetAllCars)
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carId", controllers.UpdateCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)

	return router
}
