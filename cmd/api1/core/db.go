package core

import (
	"event-driven/cmd/api1/src/domain/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	// Cambia localhost por el nombre del servicio en Docker
	dsn := "mydb:mydb@tcp(mysql:3306)/sensor?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a MySQL:", err)
		return nil, err
	}

	// Migrar la estructura de datos
	db.AutoMigrate(&entities.SensorData{})

	log.Println("Conectado a MySQL correctamente")
	return db, nil
}
