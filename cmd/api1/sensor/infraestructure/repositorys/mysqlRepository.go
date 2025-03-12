package repositorys

import (
	"api1/sensor/domain/entities"
	"gorm.io/gorm"
)

type MySQLSensorRepository struct {
	db *gorm.DB
}

// Constructor de MySQLSensorRepository
func NewMySQLSensorRepository(db *gorm.DB) *MySQLSensorRepository {
    return &MySQLSensorRepository{db: db}
}

// Implementación del método para guardar datos del sensor
func (r *MySQLSensorRepository) Save(sensorData entities.SensorData) error {
	query := `INSERT INTO sensordata (temperature, humidity) VALUES (?, ?)`

	result := r.db.Exec(query, sensorData.Temperature, sensorData.Humidity)
	return result.Error
}
