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
}

func NewSensorController(usecase *usecases.SensorUsecase) *SensorController {
	return &SensorController{
		Usecase:         usecase,
	}
}

// ReceiveSensorData maneja la recepción de datos del sensor
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}

// SendSensorData devuelve todos los datos de los sensores almacenados
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

// SubscribeToken maneja la suscripción del token a un tema en Firebase
func (c *SensorController) SubscribeToken(ctx *gin.Context) {
	log.Println("pasa aqui")
	var request struct {
		Token string `json:"token"`
	}
	log.Println("pasa aqui2")
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token requerido"})
		return
	}
	log.Println("pasa aqui3")
	// Suscribir el token al tema "all" en Firebase
	err := adapters.SubscribeToTopic(request.Token, "all")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error suscribiendo token: %v", err)})
		return
	}
	log.Println("pasa aqui4")

	ctx.JSON(http.StatusOK, gin.H{"message": "Token suscrito correctamente"})
}
