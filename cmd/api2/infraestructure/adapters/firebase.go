package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitializeFirebase() error {
	// Definimos las credenciales directamente en el código
	creds := map[string]interface{}{
		//ingresar aqui las credenciales
	}
	credsJSON, err := json.Marshal(creds)
    if err != nil {
        log.Fatalf("Error al convertir credenciales a JSON: %v", err)
    }

	// Utiliza las credenciales directamente para inicializar Firebase
	opt := option.WithCredentialsJSON(credsJSON)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error inicializando Firebase: %v", err)
	}

	App = app
	log.Println("✅ Firebase inicializado correctamente")
	return nil
}

func SendNotification(title, body string, topic string) error {
	if App == nil {
		log.Println("🚨 Firebase no está inicializado correctamente")
		return fmt.Errorf("firebase no está inicializado")
	}
	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		log.Println("🚨 Error obteniendo cliente de mensajería:", err)
		return fmt.Errorf("error obteniendo cliente de mensajería: %v", err)
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: topic,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println("🚨 Error enviando mensaje a FCM:", err)
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Notificación enviada con éxito:", response)
	return nil
}

func SubscribeToTopic(token, topic string) error {
	if App == nil {
		return fmt.Errorf("firebase no está inicializado")
	}

	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		log.Println("🚨 Error obteniendo cliente de mensajería:", err)
		return fmt.Errorf("error obteniendo cliente de mensajería: %v", err)
	}

	response, err := client.SubscribeToTopic(ctx, []string{token}, topic)
	if err != nil {
		log.Println("🚨 Error suscribiendo token al tema", topic, ":", err)
		return fmt.Errorf("error suscribiendo token al tema %s: %v", topic, err)
	}

	log.Printf("✅ Token suscrito correctamente al tema %s: %v", topic, response)
	return nil
}
