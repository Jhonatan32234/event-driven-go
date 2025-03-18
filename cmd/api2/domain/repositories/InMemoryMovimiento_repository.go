package repositories

import (
	"api2/domain/entities"
	"sync"
)

// InMemorySensorRepository es un repositorio en memoria para almacenar los datos del sensor
type InMemoryMovimientoRepository struct {
	sync.Mutex
	data []entities.MovimientoData
}

func NewInMemoryMovimientoRepository() *InMemoryMovimientoRepository {
	return &InMemoryMovimientoRepository{}
}

func (r *InMemoryMovimientoRepository) Create(movimientoData entities.MovimientoData) error {
	r.Lock()
	defer r.Unlock()
	r.data = append(r.data, movimientoData)
	return nil
}

func (r *InMemoryMovimientoRepository) GetAll() ([]entities.MovimientoData, error) {
	r.Lock()
	defer r.Unlock()
	return r.data, nil
}
