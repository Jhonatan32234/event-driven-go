package repositorys

import (
	"api1/src/domain/entities"
	"api1/src/infraestructure/adapaters"
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
