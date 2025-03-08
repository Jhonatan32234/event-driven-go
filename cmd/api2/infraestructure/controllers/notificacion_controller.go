package controllers

import (
	//"api/application/usecases"
	//"api/domain/entities"

	"event-driven/cmd/api2/application/useCases"
	"event-driven/cmd/api2/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


// Controlador para los datos del sensor
type SensorController struct {
	Usecase *useCases.SensorUsecase
}

// Método para recibir datos del sensor (POST)
func (c *SensorController) ReceiveSensorData(ctx *gin.Context) {
	var sensorData entities.SensorData
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		// Si hay un error al leer los datos, devuelve un error 400
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log para depuración
	log.Printf("Datos del sensor recibidos: %v", sensorData)

	// Almacena los datos del sensor en el repositorio
	c.Usecase.Store(sensorData)

	// Responde con un mensaje indicando que los datos se recibieron
	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}


// Método para enviar los datos del sensor (GET)
func (c *SensorController) SendSensorData(ctx *gin.Context) {
	// Recupera todos los datos almacenados
	sensorData := c.Usecase.GetAll()

	// Log para depuración
	log.Printf("Datos del sensor recuperados: %v", sensorData)

	// Devuelve los datos en formato JSON como respuesta
	ctx.JSON(http.StatusOK, gin.H{"sensorData": sensorData})
}