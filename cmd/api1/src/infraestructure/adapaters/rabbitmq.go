package adapaters

import (
	"encoding/json"
	"event-driven/cmd/api1/src/domain/entities"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQ(url, queue string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn, ch, queue}, nil
}

func (r *RabbitMQ) Publish(data entities.SensorData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.channel.Publish("", r.queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})
}

func PublishSensorEvent(data entities.SensorData) error {

	// Usar las variables de entorno para conectarse a RabbitMQ
	queue, err := NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/", "sensorQueue")
	if err != nil {
		return err
	}
	return queue.Publish(data)
}