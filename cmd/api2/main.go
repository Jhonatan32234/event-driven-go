

package main

import (
	"api2/dependencies"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	router, _ := dependencies.InitializeServer()

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno por defecto")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":"+port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
