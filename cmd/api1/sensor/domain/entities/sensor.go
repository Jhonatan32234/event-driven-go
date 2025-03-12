package entities

type SensorData struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}