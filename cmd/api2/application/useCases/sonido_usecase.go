package usecases

import (
	"api2/domain/entities"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"fmt"
	"log"
)

type SonidoUsecase struct {
	Repository repositories.SonidoRepository
}

func NewSonidoUsecase(repository repositories.SonidoRepository) *SonidoUsecase {
	return &SonidoUsecase{
		Repository: repository,
	}
}

func (uc *SonidoUsecase) Create(sonidoData entities.SonidoData) error {
	// Guardar datos del sensor
	err := uc.Repository.Create(sonidoData)
	if err != nil {
		log.Println("Error al guardar los datos del sensor de sonido:", err)
		return err
	}

	// Enviar notificación a través de Firebase
	title := "Nuevo dato del sensor de sonido"
	body := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",sonidoData.Tipo,sonidoData.Estado,sonidoData.Descripcion)	
	err = adapters.SendNotification(title,body)
	if err != nil {
		log.Println("Error enviando notificación:", err)
		return err
	}

	return nil
}

func (uc *SonidoUsecase) GetAll() ([]entities.SonidoData, error) {
	sonidoData, err := uc.Repository.GetAll()
	if err != nil {
		log.Println("Error al obtener los datos del sensor de sonido:", err)
		return nil, err
	}
	return sonidoData, nil
}