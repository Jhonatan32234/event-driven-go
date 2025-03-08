package routes

import (
	"event-driven/cmd/api1/src/application"
	"event-driven/cmd/api1/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(useCase *application.SensorUseCase) *gin.Engine {
	r := gin.Default()
	controller := controllers.NewSensorController(useCase)
	r.POST("/sensor", controller.HandleSensorData)
	return r
}
