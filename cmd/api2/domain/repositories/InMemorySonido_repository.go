package repositories

import (
	"api2/domain/entities"
	"sync"
)

// InMemorySensorRepository es un repositorio en memoria para almacenar los datos del sensor
type InMemorySonidoRepository struct {
	sync.Mutex
	data []entities.SonidoData
}

func NewInMemorySonidoRepository() *InMemorySonidoRepository {
	return &InMemorySonidoRepository{}
}

func (r *InMemorySonidoRepository) Create(sonidoData entities.SonidoData) error {
	r.Lock()
	defer r.Unlock()
	r.data = append(r.data, sonidoData)
	return nil
}

func (r *InMemorySonidoRepository) GetAll() ([]entities.SonidoData, error) {
	r.Lock()
	defer r.Unlock()
	return r.data, nil
}
