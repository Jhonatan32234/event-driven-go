package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"event-driven/cmd/api2/infraestructure/routes"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
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

	// Inicializar Firebase
	opt := option.WithCredentialsFile("/app/google-services.json") // Asegúrate de tener el archivo JSON con las credenciales de Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error inicializando Firebase: %v", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Error obteniendo cliente de Firebase: %v", err)
	}

	// Configura las rutas con Firebase
	routes.SetupRouter(router, client)

	// Inicia el servidor en el puerto 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
