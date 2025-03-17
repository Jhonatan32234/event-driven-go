package repositories

import "api2/domain/entities"

// SensorRepository define los métodos que nuestro repositorio debe implementar
type SensorRepository interface {
	Store(sensorData entities.SensorData) error
	GetAll() ([]entities.SensorData, error)
}
