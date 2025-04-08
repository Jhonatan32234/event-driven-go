package routes

import (
	"api2/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas del API
func SetupRouter(router *gin.Engine, sensorController *controllers.SensorController) {
	
		router.POST("/sensor", sensorController.ReceiveSensorData) // Recibe datos del sensor
		router.GET("/sensor", sensorController.SendSensorData)   // Envía datos del sensor
		router.POST("/subscribe", sensorController.SubscribeToken)  // Suscribe el token a Firebase
}

