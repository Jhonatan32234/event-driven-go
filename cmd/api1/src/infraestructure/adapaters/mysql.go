package adapaters

import (
	"api1/src/domain/entities"
	"log"

	"gorm.io/gorm"
)

//  se encarga de insertar los datos de SensorData en la base de datos.
func SaveDataSensor(db *gorm.DB, data entities.SensorData) error {
	// Realizar la inserción en la base de datos.
	if err := db.Create(&data).Error; err != nil {
		log.Printf("Error al insertar datos: %v", err)
		return err
	}
	log.Println("Datos insertados correctamente en la base de datos")
	log.Println(data)
	return nil
}