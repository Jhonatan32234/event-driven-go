package main

import (
	"api1/database"
	"api1/rabbit/application"
	rabbitInfra "api1/rabbit/infraestructure"
	"api1/sensor/application"
	"api1/sensor/infraestructure/repositorys"
	"api1/sensor/infraestructure/routes"
	"log"
)

func main() {
	db, err := database.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}

	// Sensor Repositorio y Servicio
	sensorRepo := repositorys.NewMySQLSensorRepository(db)
	createSensorUseCase := applications.NewCreateSensorUseCase(sensorRepo)
	sensorService := applications.NewSensorService(createSensorUseCase)

	// RabbitMQ Repositorio y Servicio
	rabbitRepo, err := rabbitInfra.NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/", "sensorQueue")
	if err != nil {
		log.Fatal(err)
	}
	publishEventUseCase := applicationr.NewPublishEventUseCase(rabbitRepo)
	rabbitService := applicationr.NewRabbitService(publishEventUseCase)

	// Router
	r := routes.SetupRouter(sensorService, rabbitService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
