package main

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/toe-royalapp/my-web-app/config"
)

func main() {
	app, _, _ := config.SetupFirebase()
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	sendAll(context.Background(), client)
}

func sendAll(ctx context.Context, client *messaging.Client) {
	// This registration token comes from the client FCM SDKs.
	registrationToken := "YOUR_REGISTRATION_TOKEN"

	// [START send_all]
	// Create a list containing up to 500 messages.
	messages := []*messaging.Message{
		{
			Notification: &messaging.Notification{
				Title: "Price drop",
				Body:  "5% off all electronics",
			},
			Token: registrationToken,
		},
		{
			Notification: &messaging.Notification{
				Title: "Price drop",
				Body:  "2% off all books",
			},
			Topic: "readers-club",
		},
	}

	br, err := client.SendAll(context.Background(), messages)
	if err != nil {
		log.Fatalln(err)
	}

	// See the BatchResponse reference documentation
	// for the contents of response.
	fmt.Printf("%d messages were sent successfully\n", br.SuccessCount)
	// [END send_all]
}
