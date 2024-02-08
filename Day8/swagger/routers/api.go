package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"swagger/controllers"
	_ "swagger/docs"
)

// @Summary Get a car by ID
// @Description Get a car from the database by ID
// @ID get-car-by-id
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} CarResponse
// @Failure 404 {object} ErrorResponse
// @Router /cars/{id} [get]

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars/:id", controllers.GetOneCars)
	router.POST("/cars", controllers.CreateCars)
	router.GET("/cars", controllers.GetAllCars)
	router.DELETE("/cars/:id", controllers.DeleteCars)
	router.PUT("/cars/:id", controllers.UpdateCars)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
