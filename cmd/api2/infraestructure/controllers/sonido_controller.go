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

type SonidoController struct {
	Usecase          *usecases.SonidoUsecase
	WebSocketAdapter *adapters.WebSocketAdapter
}

func NewSonidoController(usecase *usecases.SonidoUsecase, websocketAdapter *adapters.WebSocketAdapter) *SonidoController {
	return &SonidoController{
		Usecase:          usecase,
		WebSocketAdapter: websocketAdapter,
	}
}

// ReceiveSensorData maneja la recepción de datos del sensor
func (c *SonidoController) ReceiveSonidoData(ctx *gin.Context) {
	var sonidoData entities.SonidoData
	if err := ctx.ShouldBindJSON(&sonidoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos del sensor de sonido recibidos: %v", sonidoData)
	err := c.Usecase.Create(sonidoData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	// Enviar los datos del sensor a los WebSockets conectados
	message := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",sonidoData.Tipo,sonidoData.Estado,sonidoData.Descripcion)	
	c.WebSocketAdapter.BroadcastMessage(message)

	ctx.JSON(http.StatusOK, gin.H{"message": "Sonido data received"})
}

// SendSensorData devuelve todos los datos de los sensores almacenados
func (c *SonidoController) SendSonidoData(ctx *gin.Context) {
	sonidoData, err := c.Usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"sonidoData": sonidoData,
	})
}