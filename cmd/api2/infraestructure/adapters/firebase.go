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
		"type":                      "service_account",
		"project_id":                "event-driven-go",
		"private_key_id":            "89b8342ec3a29d86b4ce08db6f8e1543b2445989",
		"private_key":               "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDQ8hdbPZGR3VT9\nygxsnVyBad9CCyLpKDEbyEYpai7yu2JoyC63WG8zN9aemviWPCwY3DL5G5d84rjn\njLbxZ+rNGUdHEl4ZeyVVSh0XDw+fCld/8uYKvLf5GHCTxX2jxOJiewyp8jFgnBnJ\nnXG8vDp84zeBhz+DQAbwilWsDJhoVeD/RWPs3wnXtdJPt5a2LZaw6tF/JBxjJuhY\njmCX1+/y1DsJLpb4fGH2rjBgFOuYZsyptfxeXrGT/SE9ACCcGSK23Bd+WrJ7/AVy\nO8TuhWiyN1dz5YgzX/pq0zf2Jo3HEIvd0FFxtin8Zzk6rYGmgQiZxAAVw/I1tnuW\n9nMOvrmPAgMBAAECggEAAjThF0PhFivCtB1rbStqQmADJt5/cC8a+akI4j+uhQsc\nJz259JJSaYJbWcU9y9EqqNnAD6gBcD595zNLhhge8pFXEqj4fzqXZdbiOlycp+uc\nvSUSoqYxWsf6l5NPcaibuKBSo9BpdUfT8LBt3SlI6taEpXFvgJJvOe1HQ1Zf6g+9\nOeFbNSrMU70bQJQ51ECCjE8ESt7L8S+cGdL2aRQOHKfx2P74xNldJzImlUg71jpG\nP7Q/U2ame//NqH05swXizTp8TpOY68PdPutmXZjJ1DgntP8/6lrPnPXqctEE8EY+\nvebFZx0AZ4VR4etSpXQ1RISYrQrP5yKz6lrTLz/KiQKBgQDuvIIUndZ/KByALW1k\n7Q7gh4jOUJtNDW9AmCXa9BY84vITwzOBv31gCfvTelVR3x6E15pNFB+mmhax6Y4x\nntbtbX8ZSQaY6KbuDQLwxz4DKv0/Z9SFTe/QDqGtk+LDJJEu+uREQonk3ajWVk97\nGNZezVyvDVhYPIBrzHe0wip3RwKBgQDgDhjoEGi3BZ9j7lQav+6pufxBcFd2Tx07\nSaaKIx8/d7oqxbI6UJU38XeFdktyyexVvIuHLRuc6nr2AVX5Zvpwy91eSE1ebPr7\ngwSklkM1Eedzdt6Nya/MWEyKAIPfdjpV6lCo4E+f9N0DChgOLoLa8nEChzj75wKz\nx0MaTM9feQKBgQDkVDa1MSBtDRf1H315AaEw7W/SoxVlZGv7A4lxF/IM2aFddVxV\nw6dNqz7GzEG9w/+UXgCdp5l95fG+xvnQS3KUMh1VOZqttBWzp44gau7MaNk9Qzjn\nJdsuyk+ni8FdkiOpIxwejOUFl1pbZMEvolmk4hS697B+855/e1ch8nJbEQKBgG8u\nb6uQoPXZQ2vqUy/m+D6e/Q9X+P7LaX0HIi5AGXx8JBSff76ySCm3mBDRT8VJcA5n\nfnF4r9AhBx1WMlyNfk9EyrfDdykZOT5fmIk3y2flV44TeYKwh50GYAzHDqlv2KjT\nmm0CymBcuOOOObun1uVhEzUm9t8BlnSxt5mwbeM5AoGALzYmrHjKkUxQmTI4LAV5\ndjMAUBZkdulN2ZnMcMxjxTShEiW1zgjG8XDrQE5+DlypBIoVnYcJF8x4Mw7hrpZb\n9LXX+EYWfDoI62ZfJQMB2CbMQGRKfswRR3/4XtqeX8tms4RpwRFOSw+Af0RLu02S\nhg+MopGeef3xmZe8t9WW/qM=\n-----END PRIVATE KEY-----\n",
		"client_email":              "firebase-adminsdk-fbsvc@event-driven-go.iam.gserviceaccount.com",
		"client_id":                 "111530347057227459152",
		"auth_uri":                 "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":      "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-fbsvc%40event-driven-go.iam.gserviceaccount.com",
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
