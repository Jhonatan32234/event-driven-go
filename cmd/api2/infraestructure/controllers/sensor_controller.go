package controllers

import (
	"api2/application/usecases"
	"api2/domain/entities"
	"api2/infraestructure/adapters"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	Usecase         *usecases.SensorUsecase
	WebSocketAdapter *adapters.WebSocketAdapter
}

func NewSensorController(usecase *usecases.SensorUsecase, websocketAdapter *adapters.WebSocketAdapter) *SensorController {
	return &SensorController{
		Usecase:         usecase,
		WebSocketAdapter: websocketAdapter,
	}
}

func (c *SensorController) ReceiveSensorData(ctx *gin.Context) {
	var sensorData entities.SensorData
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos del sensor recibidos: %v", sensorData)
	err := c.Usecase.Store(sensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	// Enviar los datos del sensor a los WebSockets conectados
	message := fmt.Sprintf(`{"message": "Nuevo dato del sensor:","temperature":"%.2f","humidity":"%.2f"}`, sensorData.Temperature, sensorData.Humidity)
    c.WebSocketAdapter.BroadcastMessage(message)



	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}

func (c *SensorController) SendSensorData(ctx *gin.Context) {
	sensorData, err := c.Usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sensorData": sensorData,
	})
}
