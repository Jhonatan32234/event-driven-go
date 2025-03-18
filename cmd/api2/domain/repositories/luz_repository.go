package repositories

import "api2/domain/entities"

type LuzRepository interface {
	Create(luzData entities.LuzData) error
	GetAll() ([]entities.LuzData,error)
}