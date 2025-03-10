package controllers

import (
	"api2/application/useCases"
	"api2/domain/entities"
	"api2/infraestructure/adapters"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SensorController struct {
	Usecase        *useCases.SensorUsecase
	FirebaseClient *adapters.FirebaseAdapter
}

func (c *SensorController) ReceiveSensorData(ctx *gin.Context) {
	var sensorData entities.SensorData
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos del sensor recibidos: %v", sensorData)
	c.Usecase.Store(sensorData)

	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}

func (c *SensorController) SendSensorData(ctx *gin.Context) {
	sensorData := c.Usecase.GetAll()
	log.Printf("Datos del sensor recuperados: %v", sensorData)

	// Lista de tokens simulada (estos deberían obtenerse de la base de datos)
	tokens := []string{"AIzaSyBT2zE5VRCiie6bsA2kEZBAkaieECpeLWM"}

	for _, token := range tokens {
		err := c.FirebaseClient.SendNotification(
			"Nuevos datos del sensor",
			"Se han registrado nuevos valores de temperatura y humedad.",
			token,
		)

		if err != nil {
			log.Printf("Error enviando notificación a %s: %v", token, err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sensorData": sensorData,
		"message":    "Datos enviados y notificaciones disparadas",
	})
}
