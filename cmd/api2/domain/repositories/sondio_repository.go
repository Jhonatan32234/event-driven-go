package repositories

import "api2/domain/entities"

type SonidoRepository interface {
	Create(sonidoData entities.SonidoData) error
	GetAll()([]entities.SonidoData,error)
}

