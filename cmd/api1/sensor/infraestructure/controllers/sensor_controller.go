package controllers

import (
	"api1/rabbit/application"
	applications "api1/sensor/application"
	"api1/sensor/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	sensorService *applications.SensorService
	rabbitService *applicationr.RabbitService
}

func NewSensorController(sensorService *applications.SensorService, rabbitService *applicationr.RabbitService) *SensorController {
	return &SensorController{
		sensorService: sensorService,
		rabbitService: rabbitService,
	}
}

func (s *SensorController) Execute(c *gin.Context) {
	var data entities.SensorData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar en la base de datos
	if err := s.sensorService.SaveSensorData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el sensor"})
		return
	}

	// Publicar en RabbitMQ
	if err := s.rabbitService.PublishSensorData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error publicando el evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos procesados y publicados correctamente"})
}
