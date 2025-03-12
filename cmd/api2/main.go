package main

import (
	"api2/application/usecases"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"api2/infraestructure/controllers"
	"api2/infraestructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear instancia del router Gin
	router := gin.Default()

	// Crear instancia del WebSocketAdapter
	webSocketAdapter := adapters.NewWebSocketAdapter()

	// Crear instancia del repositorio (aquí usaremos un repositorio en memoria)
	sensorRepository := repositories.NewInMemorySensorRepository()

	// Crear el caso de uso para manejar la lógica de negocio
	sensorUsecase := usecases.NewSensorUsecase(sensorRepository)

	// Crear el controlador
	sensorController := controllers.NewSensorController(sensorUsecase, webSocketAdapter)

	// Configurar las rutas
	routes.SetupRouter(router, sensorController)

	// Iniciar WebSocket en un goroutine
	go webSocketAdapter.Start()

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		conn, err := webSocketAdapter.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error al establecer WebSocket:", err)
			return
		}
		webSocketAdapter.HandleWebSocketConnection(conn)
	})

	// Iniciar el servidor en el puerto 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
