package repositories

import (
	"api2/domain/entities"
	"sync"
)

// InMemorySensorRepository es un repositorio en memoria para almacenar los datos del sensor
type InMemorySensorRepository struct {
	sync.Mutex
	data []entities.SensorData
}

func NewInMemorySensorRepository() *InMemorySensorRepository {
	return &InMemorySensorRepository{}
}

func (r *InMemorySensorRepository) Store(sensorData entities.SensorData) error {
	r.Lock()
	defer r.Unlock()
	r.data = append(r.data, sensorData)
	return nil
}

func (r *InMemorySensorRepository) GetAll() ([]entities.SensorData, error) {
	r.Lock()
	defer r.Unlock()
	return r.data, nil
}
