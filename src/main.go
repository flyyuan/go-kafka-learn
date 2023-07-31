package main

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "myTopic",
		Balancer: &kafka.LeastBytes{},
	})

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("Hello, Kafka"),
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}
}
