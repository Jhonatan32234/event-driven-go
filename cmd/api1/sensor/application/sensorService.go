package applications

import (
	"api1/sensor/domain/entities"
)

type SensorService struct {
	createSensorUseCase *CreateSensorUseCase
}

func NewSensorService(createSensorUseCase *CreateSensorUseCase) *SensorService {
	return &SensorService{createSensorUseCase: createSensorUseCase}
}

func (s *SensorService) SaveSensorData(data entities.SensorData) error {
	return s.createSensorUseCase.Execute(data)
}
