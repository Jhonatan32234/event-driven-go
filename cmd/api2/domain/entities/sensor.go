package entities

type SensorData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Emmiter     string `json:"emmiter"`
	Topic       string `json:"topic"`
	CreatedAt   string `json:"created_at"`
}
