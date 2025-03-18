package repositories

import (
	"api2/domain/entities"
	"sync"
)

// InMemorySensorRepository es un repositorio en memoria para almacenar los datos del sensor
type InMemoryLuzRepository struct {
	sync.Mutex
	data []entities.LuzData
}

func NewInMemoryLuzRepository() *InMemoryLuzRepository {
	return &InMemoryLuzRepository{}
}

func (r *InMemoryLuzRepository) Create(luzData entities.LuzData) error {
	r.Lock()
	defer r.Unlock()
	r.data = append(r.data, luzData)
	return nil
}

func (r *InMemoryLuzRepository) GetAll() ([]entities.LuzData, error) {
	r.Lock()
	defer r.Unlock()
	return r.data, nil
}
