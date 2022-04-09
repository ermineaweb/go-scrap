package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type kafkaProducer struct {
	brokerAddress string
	topic         string
}

func NewKafkaProducer(brokerAddress, topic string) *kafkaProducer {
	return &kafkaProducer{brokerAddress: brokerAddress, topic: topic}
}

// kafka producer send message to a kafka broker, on a topic
func (k *kafkaProducer) Produce(message []byte) {
	ctx := context.Background()

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{k.brokerAddress},
		Topic:   k.topic,
		Logger:  l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte("mykey"),
		Value: message,
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("send message", string(message))
}
