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

type MovimientoController struct {
	Usecase          *usecases.MovimientoUsecase
	WebSocketAdapter *adapters.WebSocketAdapter
}

func NewMovimientoController(usecase *usecases.MovimientoUsecase, websocketAdapter *adapters.WebSocketAdapter) *MovimientoController {
	return &MovimientoController{
		Usecase:          usecase,
		WebSocketAdapter: websocketAdapter,
	}
}

// ReceiveSensorData maneja la recepción de datos del sensor
func (c *MovimientoController) ReceiveMovimientoData(ctx *gin.Context) {
	var movimientoData entities.MovimientoData
	if err := ctx.ShouldBindJSON(&movimientoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos del sensor de movimiento recibidos: %v", movimientoData)
	err := c.Usecase.Create(movimientoData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	// Enviar los datos del sensor a los WebSockets conectados
	message := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",movimientoData.Tipo,movimientoData.Estado,movimientoData.Descripcion)	
	c.WebSocketAdapter.BroadcastMessage(message)

	ctx.JSON(http.StatusOK, gin.H{"message": "Movimiento data received"})
}

// SendSensorData devuelve todos los datos de los sensores almacenados
func (c *MovimientoController) SendMovimientoData(ctx *gin.Context) {
	movimientoData, err := c.Usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"movimientoData": movimientoData,
	})
}