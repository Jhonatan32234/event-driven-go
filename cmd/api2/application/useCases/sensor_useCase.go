package useCases

import (
	"event-driven/cmd/api2/domain"
	"event-driven/cmd/api2/domain/entities"
)

// Caso de uso para manejar los datos del sensor
type SensorUsecase struct {
	Repo *domain.InMemorySensorRepository
}

// Método para almacenar los datos del sensor
func (u *SensorUsecase) Store(sensorData entities.SensorData) {
	u.Repo.Store(sensorData)
}

// Método para obtener todos los datos del sensor
func (u *SensorUsecase) GetAll() []entities.SensorData {
	return u.Repo.GetAll()
}