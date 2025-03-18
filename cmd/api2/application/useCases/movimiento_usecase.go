package usecases

import (
	"api2/domain/entities"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"fmt"
	"log"
)

type MovimientoUsecase struct {
	Repository repositories.MovimientoRepository
}

func NewMovimientoUsecase(repository repositories.MovimientoRepository) *MovimientoUsecase {
	return &MovimientoUsecase{
		Repository: repository,
	}
}

func (uc *MovimientoUsecase) Create(movimientoData entities.MovimientoData) error {
	// Guardar datos del sensor
	err := uc.Repository.Create(movimientoData)
	if err != nil {
		log.Println("Error al guardar los datos del sensor de movimiento:", err)
		return err
	}

	// Enviar notificación a través de Firebase
	title := "Nuevo dato del sensor de movimiento"
	body := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",movimientoData.Tipo,movimientoData.Estado,movimientoData.Descripcion)	
	err = adapters.SendNotification(title, body)
	if err != nil {
		log.Println("Error enviando notificación:", err)
		return err
	}

	return nil
}

func (uc *MovimientoUsecase) GetAll() ([]entities.MovimientoData, error) {
	// Obtener todos los datos de los sensores
	movimientoData, err := uc.Repository.GetAll()
	if err != nil {
		log.Println("Error al obtener los datos del sensor de movimiento:", err)
		return nil, err
	}
	return movimientoData, nil
}