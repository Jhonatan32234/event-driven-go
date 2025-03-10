package main

import (
	"bytes"
	"consumer/entities"
	"encoding/json"
	"log"
	"net/http"

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

func (r *RabbitMQ) ConsumeMessages(apiURL, newQueue string) error {
	msgs, err := r.channel.Consume(r.queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	// Declarar la nueva cola antes de enviar mensajes a ella
	_, err = r.channel.QueueDeclare(newQueue, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar la nueva cola: %v", err)
	}

	log.Println("[*] Esperando mensajes. Para salir presiona CTRL+C")

	for msg := range msgs {
		var data entities.SensorData
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			log.Printf("Error al deserializar mensaje: %v", err)
			continue
		}
		log.Println("Datos recibidos:", data)
		sendToAPI(apiURL, data)
		sendToQueue(r, newQueue, data)
	}
	return nil
}

func sendToAPI(apiURL string, data entities.SensorData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializando JSON: %v", err)
		return
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error enviando datos a API: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Datos enviados a API con estado: %s", resp.Status)
}

func sendToQueue(r *RabbitMQ, queueName string, data entities.SensorData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializando JSON para la cola: %v", err)
		return
	}

	err = r.channel.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})
	if err != nil {
		log.Printf("Error enviando datos a la cola: %v", err)
		return
	}

	log.Printf("Datos enviados a la cola %s correctamente", queueName)
}

func main() {

	rmq, err := NewRabbitMQ("amqp://guest:guest@rabbitmq:5672/", "sensorQueue")
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	defer rmq.conn.Close()
	defer rmq.channel.Close()

	if err := rmq.ConsumeMessages("http://api2:8000/receive", "NEW_QUEUE"); err != nil {
		log.Fatalf("Error consumiendo mensajes: %v", err)
	}
}
