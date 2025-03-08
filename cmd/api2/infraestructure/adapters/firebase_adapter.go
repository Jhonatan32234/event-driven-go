package adapters

import (
	"context"
	"firebase.google.com/go/messaging"
	"log"
)

type FirebaseAdapter struct {
	client *messaging.Client
}

func NewFirebaseAdapter(client *messaging.Client) *FirebaseAdapter {
	return &FirebaseAdapter{client: client}
}

func (fa *FirebaseAdapter) SendNotification(title, message, token string) error {
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  message,
		},
		Token: token,
	}

	_, err := fa.client.Send(context.Background(), msg)
	if err != nil {
		log.Printf("Error enviando notificación: %v", err)
		return err
	}

	log.Println("Notificación enviada correctamente")
	return nil
}
