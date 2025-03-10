package routes

import (
	"api2/application/useCases"
	"api2/domain"
	"api2/infraestructure/adapters"
	"api2/infraestructure/controllers"
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
