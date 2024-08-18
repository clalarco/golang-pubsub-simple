# Simple example for Pubsub

It uses the Google Cloud pubsub emulator.

Check instructions at [Google Cloud documentation](https://cloud.google.com/pubsub/docs/emulator).

Start the emulator:
```bash
gcloud beta emulators pubsub start --project=pubsub-test
```

Run publisher, and create topic if it doesn't exist:
```bash
go run ./publisher/publisher.go
```

Run subscriber, which waits for the messages:
```bash
go run ./subscriber/subscriber.go
```
