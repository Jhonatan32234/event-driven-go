package applicationr

import "api1/sensor/domain/entities"

type RabbitService struct {
	publishEventUseCase *PublishEventUseCase
}

func NewRabbitService(publishEventUseCase *PublishEventUseCase) *RabbitService {
	return &RabbitService{publishEventUseCase: publishEventUseCase}
}

func (r *RabbitService) PublishSensorData(data entities.SensorData) error {
	return r.publishEventUseCase.Execute(data)
}
