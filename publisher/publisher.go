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

	// Sets your Google Cloud Platform project ID.
	projectID := "pubsub-test"
	topicID := "topic-01"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	topic, err := GetTopic(ctx, client, topicID)
	if err != nil {
		println(err.Error())
		println("Error on init")
		os.Exit(1)
	}
	result := topic.Publish(ctx, &pubsub.Message{Data: []byte("Ola ke ase!")})
	serverID, error := result.Get(ctx)
	fmt.Printf("Published message: %v, error: %v\n", serverID, error)
	os.Exit(0)
}

func GetTopic(ctx context.Context, client *pubsub.Client, topicID string) (*pubsub.Topic, error) {
	// Check if the topic exists
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to check if topic exists: %v", err)
		return nil, err
	}

	if !exists {
		// Create the topic if it doesn't exist
		topic, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
			return nil, err
		}
		fmt.Printf("Topic %s created successfully\n", topicID)
	} else {
		fmt.Printf("Topic %s already exists\n", topicID)
	}
	return topic, nil
}
