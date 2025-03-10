package main

import (
	"api1/core"
	"api1/src/application"
	"api1/src/domain"
	"api1/src/infraestructure/routes"
	"log"
)

func main() {
	db, err := core.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}

	repo := domain.NewSensorRepositoryDB(db)
	useCase := application.NewSensorUseCase(repo)
	r := routes.SetupRouter(useCase)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err) 
	}
}
