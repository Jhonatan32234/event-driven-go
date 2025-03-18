package dependencies

import (
	"api2/infraestructure/adapters"
	
	"api2/domain/repositories"
	"api2/application/usecases"
	"api2/infraestructure/controllers"
	"api2/infraestructure/routes"
	"log"

	"github.com/gin-gonic/gin"
	"net/http"
)

// Configuración de CORS
func CorsMiddleware() gin.HandlerFunc {
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

// Inicializar y devolver las dependencias del servidor
func InitializeServer() (*gin.Engine, *adapters.WebSocketAdapter, *controllers.SensorController) {
	// Inicializar Firebase
	err := adapters.InitializeFirebase()
	if err != nil {
		log.Fatalf("Error al inicializar Firebase: %v", err)
	}

	// Crear instancia del router Gin
	router := gin.Default()

	// Aplicar el middleware CORS antes de configurar las rutas
	router.Use(CorsMiddleware())

	// Crear instancia del WebSocketAdapter
	webSocketAdapter := adapters.NewWebSocketAdapter()

	// Crear instancia del repositorio (aquí usaremos un repositorio en memoria)
	sensorRepository := repositories.NewInMemorySensorRepository()
	luzRepository := repositories.NewInMemoryLuzRepository()
	sonidoRepository := repositories.NewInMemorySonidoRepository()
	movimientoRepository := repositories.NewInMemoryMovimientoRepository()


	// Crear el caso de uso para manejar la lógica de negocio
	sensorUsecase := usecases.NewSensorUsecase(sensorRepository)
	luzUsecase := usecases.NewLuzUsecase(luzRepository)
	sonidousecase := usecases.NewSonidoUsecase(sonidoRepository)
	movimientoUsecase := usecases.NewMovimientoUsecase(movimientoRepository)


	// Crear el controlador
	sensorController := controllers.NewSensorController(sensorUsecase, webSocketAdapter)
	luzController := controllers.NewLuzController(luzUsecase,webSocketAdapter)
	sonidoController := controllers.NewSonidoController(sonidousecase,webSocketAdapter)
	movimientoController := controllers.NewMovimientoController(movimientoUsecase,webSocketAdapter)


	// Configurar las rutas
	routes.SetupRouter(router, sensorController,luzController,sonidoController,movimientoController)

	return router, webSocketAdapter, sensorController
}
