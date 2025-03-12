package routes

import (
	applicationr "api1/rabbit/application"
	applications "api1/sensor/application"
	"api1/sensor/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(sensorService *applications.SensorService, rabbitService *applicationr.RabbitService) *gin.Engine {
	r := gin.Default()
	controller := controllers.NewSensorController(sensorService, rabbitService)
	r.POST("/sensor", controller.Execute)
	return r
}
