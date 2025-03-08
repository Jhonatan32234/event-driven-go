package routes

import (
	"event-driven/cmd/api2/application/useCases"
	"event-driven/cmd/api2/domain"
	"event-driven/cmd/api2/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

// Configuración de las rutas
func SetupRouter(router *gin.Engine) {
	// Creación del repositorio y del caso de uso
	sensorRepo := &domain.InMemorySensorRepository{}
	sensorUsecase := &useCases.SensorUsecase{Repo: sensorRepo}
	sensorController := &controllers.SensorController{Usecase: sensorUsecase}

	// Rutas para recibir y enviar datos del sensor
	router.POST("/receive", sensorController.ReceiveSensorData)
	router.GET("/send-data", sensorController.SendSensorData)
}
