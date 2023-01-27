package produce_service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"net/http"
)

func SendMsg(w http.ResponseWriter, r *http.Request) {
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.StringEncoder("Hello world"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent to partition %d at offset %d  value = %v \n", partition, offset, msg.Value)
}
