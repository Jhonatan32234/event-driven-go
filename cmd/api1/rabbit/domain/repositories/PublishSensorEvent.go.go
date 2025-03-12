package repositories

import "api1/sensor/domain/entities"

type PublishSensorEventRepository interface {
	Publish(data entities.SensorData) error
}

