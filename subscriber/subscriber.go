package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	if os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085") != nil {
		os.Exit(1)
	}

	ctx := context.Background()

	projectID := "pubsub-test"
	subscriptionID := "subscription-01"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	subscription, err := GetSubscription(ctx, client, subscriptionID)
	if err != nil {
		println(err.Error())
	}

	for {
		err := subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
			fmt.Printf("Received message: %s\n", m.Data)
			m.Ack() // Acknowledge that we've consumed the message.
		})
		if err != nil {
			println(err.Error())
		}
	}
}

func GetSubscription(ctx context.Context, client *pubsub.Client, subscriptionID string) (*pubsub.Subscription, error) {
	// Check if the topic exists
	subscription := client.Subscription(subscriptionID)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if subscription exists: %v", err)
		return nil, err
	}

	if !exists {
		cfg := pubsub.SubscriptionConfig{
			Topic: client.Topic("topic-01"),
		}
		// Create the subscription if it doesn't exist
		subscription, err = client.CreateSubscription(ctx, subscriptionID, cfg)
		if err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
			return nil, err
		}
		fmt.Printf("Subscription %s created successfully\n", subscriptionID)
	} else {
		fmt.Printf("Using existing subscription %s\n", subscriptionID)
	}
	return subscription, nil
}
