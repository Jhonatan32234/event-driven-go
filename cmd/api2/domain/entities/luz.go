package entities

type LuzData struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Tipo string `json:"tipo"`
	Estado string `json:"estado"`
	Descripcion string `json:"descripcion"`
}
