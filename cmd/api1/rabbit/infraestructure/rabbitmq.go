package infraestructure

import (
	"api1/rabbit/domain/repositories"
	"api1/sensor/domain/entities"
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

var _ repositories.PublishSensorEventRepository = (*RabbitMQ)(nil)

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
