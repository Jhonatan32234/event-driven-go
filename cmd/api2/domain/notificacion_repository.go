package domain

import (
	//"api/domain/entities"
	"event-driven/cmd/api2/domain/entities"
	"log"
	"sync"
)

// Repositorio en memoria que almacena los datos del sensor
type InMemorySensorRepository struct {
	sensorData []entities.SensorData
	mu         sync.Mutex
}

// Método para almacenar los datos del sensor
func (r *InMemorySensorRepository) Store(sensorData entities.SensorData) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sensorData = append(r.sensorData, sensorData)

	// Log para depuración
	log.Printf("Datos del sensor almacenados: %v", sensorData)
}

// Método para obtener todos los datos del sensor
func (r *InMemorySensorRepository) GetAll() []entities.SensorData {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.sensorData
}