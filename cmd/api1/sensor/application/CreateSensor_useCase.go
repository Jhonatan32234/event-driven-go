package applications

import (
	"api1/sensor/domain"
	"api1/sensor/domain/entities"
)

type CreateSensorUseCase struct {
	repo domain.SensorRepository
}

func NewCreateSensorUseCase(repo domain.SensorRepository) *CreateSensorUseCase {
	return &CreateSensorUseCase{repo: repo}
}

func (s *CreateSensorUseCase) Execute(data entities.SensorData) error {
	return s.repo.Save(data)
}
