package entities

type SonidoData struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Tipo string `json:"tipo"`
	Estado string `json:"estado"`
	Descripcion string `json:"descripcion"`
}
