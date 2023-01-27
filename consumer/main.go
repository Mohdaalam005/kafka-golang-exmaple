package main

import (
	"github.com/Shopify/sarama"
	"log"
	"net/http"
)

func main() {

	log.Println(1)
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V2_4_0_0

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the topic
	partitionConsumer, err := consumer.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Error subscribing to topic: %v", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: key=%s, value=%s", string(msg.Key), string(msg.Value))
	}

	http.ListenAndServe(":9090", nil)
}
