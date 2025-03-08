package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"event-driven/cmd/api2/infraestructure/routes"
)

func main() {
	// Crea una nueva instancia de Gin
	router := gin.Default()

	// Aplica CORS **antes** de configurar las rutas
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configura las rutas, pasando el `router` ya creado
	routes.SetupRouter(router)

	// Inicia el servidor en el puerto 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
