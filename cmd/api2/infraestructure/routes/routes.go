package routes

import (
	"api2/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas del API
func SetupRouter(router *gin.Engine, sensorController *controllers.SensorController,luzController *controllers.LuzController,sonidoController *controllers.SonidoController,movimientoController *controllers.MovimientoController) {
	
		router.POST("/sensor", sensorController.ReceiveSensorData) // Recibe datos del sensor
		router.GET("/sensor", sensorController.SendSensorData)   // Envía datos del sensor
		router.POST("/subscribe", sensorController.SubscribeToken)  // Suscribe el token a Firebase
		router.POST("/luz",luzController.ReceiveLuzData)
		router.GET("/luz",luzController.SendLuzData)
		router.POST("/sonido",sonidoController.ReceiveSonidoData)
		router.GET("/sonido",sonidoController.SendSonidoData)
		router.POST("/movimiento",movimientoController.ReceiveMovimientoData)
		router.GET("/movimiento",movimientoController.SendMovimientoData)
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