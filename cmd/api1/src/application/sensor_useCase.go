package application

import (
	"api1/src/domain"
	"api1/src/domain/entities"
)

type SensorUseCase struct {
	repo domain.SensorRepository
}

func NewSensorUseCase(repo domain.SensorRepository) *SensorUseCase {
	return &SensorUseCase{repo: repo}
}

func (s *SensorUseCase) ProcessSensorData(data entities.SensorData) error {
	return s.repo.Save(data)
}
