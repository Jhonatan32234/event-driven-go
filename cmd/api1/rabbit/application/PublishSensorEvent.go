package applicationr

import (
	"api1/sensor/domain/entities"
	"api1/rabbit/domain/repositories"
)

type PublishEventUseCase struct {
	eventPublisher repositories.PublishSensorEventRepository
}

func NewPublishEventUseCase(eventPublisher repositories.PublishSensorEventRepository) *PublishEventUseCase {
	return &PublishEventUseCase{eventPublisher: eventPublisher}
}

func (p *PublishEventUseCase) Execute(data entities.SensorData) error {
	return p.eventPublisher.Publish(data)
}
