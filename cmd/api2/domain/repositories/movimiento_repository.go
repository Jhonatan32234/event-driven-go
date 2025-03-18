package repositories

import "api2/domain/entities"

type MovimientoRepository interface {
	Create(movimientoData entities.MovimientoData) error
	GetAll() ([]entities.MovimientoData,error)
}