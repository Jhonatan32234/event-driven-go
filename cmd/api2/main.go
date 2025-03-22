/*package main

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
*/

/*

package main

import (
	"api2/application/usecases"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"api2/infraestructure/controllers"
	"api2/infraestructure/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware CORS personalizado
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Configurar los encabezados CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Manejar solicitudes OPTIONS
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continuar con el siguiente middleware o handler
		c.Next()
	}
}

func main() {
	// Inicializar Firebase
	err := adapters.InitializeFirebase()
	if err != nil {
		log.Fatalf("🚨 Error al inicializar Firebase: %v", err)
	}

	// Crear instancia del router Gin
	router := gin.Default()

	// Aplicar el middleware CORS antes de configurar las rutas
	router.Use(corsMiddleware())

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
*/

package main

import (
	"api2/dependencies"
	"log"

)

func main() {
	// Inicializar el servidor, WebSocket y controlador
	router, _ := dependencies.InitializeServer()

	// Iniciar WebSocket en un goroutine
	

	// WebSocket route

	// Iniciar el servidor en el puerto 8000
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
