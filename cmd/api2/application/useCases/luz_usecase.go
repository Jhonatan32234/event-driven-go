package usecases

import (
	"api2/domain/entities"
	"api2/domain/repositories"
	"api2/infraestructure/adapters"
	"fmt"
	"log"
)

type LuzUsecase struct {
	Repository repositories.LuzRepository
}

func NewLuzUsecase(repostory repositories.LuzRepository) *LuzUsecase {
	return &LuzUsecase{
		Repository: repostory,
	}
}

func (uc *LuzUsecase) Create(luzData entities.LuzData) error{
	err := uc.Repository.Create(luzData)
	if err != nil {
		log.Println("Error al guardar los datos del sensor:", err)
		return err
	}
	log.Print("llama aqui")
	title := "Nuevo dato del sensor de Luz"
	body := fmt.Sprintf("Tipo: %s, Estado: %s, Descripcion: %s",luzData.Tipo,luzData.Estado,luzData.Descripcion)
	err = adapters.SendNotification(title,body)
	if err != nil {
		log.Println("Error enviando notificación:", err)
		return err
	}

	return nil
}


func (uc *LuzUsecase) GetAll() ([]entities.LuzData,error){
	luzData, err := uc.Repository.GetAll()
	if err != nil {
		log.Println("Error al obtener los datos del sensor:", err)
		return nil, err
	}
	return luzData, nil
}
