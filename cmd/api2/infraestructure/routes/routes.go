package routes

import (
	"event-driven/cmd/api2/application/useCases"
	"event-driven/cmd/api2/domain"
	"event-driven/cmd/api2/infraestructure/adapters"
	"event-driven/cmd/api2/infraestructure/controllers"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, firebaseClient *messaging.Client) {
	sensorRepo := &domain.InMemorySensorRepository{}
	sensorUsecase := &useCases.SensorUsecase{Repo: sensorRepo}
	firebaseAdapter := adapters.NewFirebaseAdapter(firebaseClient)
	sensorController := &controllers.SensorController{
		Usecase:        sensorUsecase,
		FirebaseClient: firebaseAdapter,
	}

	router.POST("/receive", sensorController.ReceiveSensorData)
	router.GET("/send-data", sensorController.SendSensorData)
}
