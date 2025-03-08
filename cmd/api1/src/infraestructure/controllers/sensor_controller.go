package controllers

import (
	"event-driven/cmd/api1/src/application"
	"event-driven/cmd/api1/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	useCase *application.SensorUseCase
}

func NewSensorController(useCase *application.SensorUseCase) *SensorController {
	return &SensorController{useCase: useCase}
}

func (s *SensorController) HandleSensorData(c *gin.Context) {
	var data entities.SensorData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.useCase.ProcessSensorData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Datos procesados correctamente"})
}