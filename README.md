# Simple example for Pubsub

It uses the Google Cloud pubsub emulator.

Check instructions at [Google Cloud documentation](https://cloud.google.com/pubsub/docs/emulator).

Run publisher, and create topic if it doesn't exist:
```bash
go run ./publisher/publisher.go
```

Run subscriber, which waits for the messages:
```bash
go run ./subscriber/subscriber.go
```
