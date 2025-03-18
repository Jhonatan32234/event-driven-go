package usecases

import (
	"api2/domain/entities"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"fmt"
	"log"
)

type SensorUsecase struct {
	Repository repositories.SensorRepository
}

func NewSensorUsecase(repository repositories.SensorRepository) *SensorUsecase {
	return &SensorUsecase{
		Repository: repository,
	}
}

func (uc *SensorUsecase) Store(sensorData entities.SensorData) error {
	// Guardar datos del sensor
	err := uc.Repository.Store(sensorData)
	if err != nil {
		log.Println("Error al guardar los datos del sensor:", err)
		return err
	}

	// Enviar notificación a través de Firebase
	title := "Nuevo dato del sensor"
	body := fmt.Sprintf("Temperatura: %.2f, Humedad: %.2f", sensorData.Temperature, sensorData.Humidity)
	err = adapters.SendNotification(title, body)
	if err != nil {
		log.Println("Error enviando notificación:", err)
		return err
	}

	return nil
}

func (uc *SensorUsecase) GetAll() ([]entities.SensorData, error) {
	// Obtener todos los datos de los sensores
	sensorData, err := uc.Repository.GetAll()
	if err != nil {
		log.Println("Error al obtener los datos del sensor:", err)
		return nil, err
	}
	return sensorData, nil
}