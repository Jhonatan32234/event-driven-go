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

type LuzController struct {
	Usecase          *usecases.LuzUsecase
	WebSocketAdapter *adapters.WebSocketAdapter
}

func NewLuzController(usecase *usecases.LuzUsecase, websocketAdapter *adapters.WebSocketAdapter) *LuzController {
	return &LuzController{
		Usecase:          usecase,
		WebSocketAdapter: websocketAdapter,
	}
}

func (c *LuzController) ReceiveLuzData(ctx *gin.Context) {
	log.Println("pasa aqui1")
	var luzData entities.LuzData
	if err := ctx.ShouldBindJSON(&luzData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Datos del sensor de luz recibidos: %v", luzData)
	err := c.Usecase.Create(luzData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	// Enviar los datos del sensor a los WebSockets conectados
	message := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",luzData.Tipo,luzData.Estado,luzData.Descripcion)	
	c.WebSocketAdapter.BroadcastMessage(message)

	ctx.JSON(http.StatusOK, gin.H{"message": "Luz data received"})
}

func (c *LuzController) SendLuzData(ctx *gin.Context) {
	luzData, err := c.Usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"luzData": luzData,
	})
}