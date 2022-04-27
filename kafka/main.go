package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "movie",
	}
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("movies"),
	})
	if err != nil {
		log.Fatalf("cannot write a messager: %v", err)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "movie",
		Topic:    "movie",
		MinBytes: 0,
		MaxBytes: 10e6,
	})
	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatalf("cannot read a message: %v", err)
			reader.Close()
		}
		fmt.Print("Received a message:", string(message.Value))
	}
	reader.Close()
}
