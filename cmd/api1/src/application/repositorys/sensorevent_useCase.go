package repositorys

import (
	"event-driven/cmd/api1/src/domain/entities"
	"event-driven/cmd/api1/src/infraestructure/adapaters"
)



type PublishSensorEventRepository interface {
	Publish(data entities.SensorData) error
}

type PublishSensorEventRepositoryImpl struct {}

func NewPublishSensorEventRepository() *PublishSensorEventRepositoryImpl {
	return &PublishSensorEventRepositoryImpl{}
}

func (r *PublishSensorEventRepositoryImpl) Publish(data entities.SensorData) error {
	return adapaters.PublishSensorEvent(data)
}
