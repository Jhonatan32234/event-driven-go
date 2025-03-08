package domain

import (
	
	"event-driven/cmd/api1/src/domain/entities"
	"event-driven/cmd/api1/src/infraestructure/adapaters"

	"gorm.io/gorm"
)

type SensorRepository interface {
	Save(data entities.SensorData) error
}

type SensorRepositoryDB struct {
	db *gorm.DB
}

func NewSensorRepositoryDB(db *gorm.DB) *SensorRepositoryDB {
	return &SensorRepositoryDB{db: db}
}

func (r *SensorRepositoryDB) Save(data entities.SensorData) error {
	if err := adapaters.SaveDataSensor(r.db, data); err != nil {
		return err
	}
	return adapaters.PublishSensorEvent(data)
}
