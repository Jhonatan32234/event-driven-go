package adapters

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitializeFirebase() error {
	credsJSON := os.Getenv("FIREBASE_CREDENTIALS_JSON")
	if credsJSON == "" {
		return fmt.Errorf("FIREBASE_CREDENTIALS_JSON no está definida")
	}

	opt := option.WithCredentialsJSON([]byte(credsJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error inicializando Firebase: %v", err)
	}

	App = app
	log.Println("✅ Firebase inicializado correctamente")
	return nil
}

func SendNotification(title, body string) error {
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
		Topic: "all",
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println("🚨 Error enviando mensaje a FCM:", err)
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Notificación enviada con éxito:", response)
	return nil
}

// SubscribeToTopic suscribe un token a un tema en Firebase
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
