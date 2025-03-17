package routes

import (
	"api2/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas del API
func SetupRouter(router *gin.Engine, sensorController *controllers.SensorController) {
	api := router.Group("/api")
	{
		api.POST("/receive", sensorController.ReceiveSensorData) // Recibe datos del sensor
		api.GET("/send-data", sensorController.SendSensorData)   // Envía datos del sensor
		api.POST("/subscribe", sensorController.SubscribeToken)  // Suscribe el token a Firebase
	}
}


/*package routes

import (
	"api2/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas del API
func SetupRouter(router *gin.Engine, sensorController *controllers.SensorController) {
	api := router.Group("/api")
	{
		api.POST("/receive", sensorController.ReceiveSensorData) // Recibe datos del sensor
		api.GET("/send-data", sensorController.SendSensorData)   // Envía datos del sensor
	}
}
*/