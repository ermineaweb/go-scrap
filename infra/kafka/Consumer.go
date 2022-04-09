package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type kafkaConsumer struct {
	brokerAddress string
	topic         string
}

func NewKafkaConsumer(brokerAddress, topic string) *kafkaConsumer {
	return &kafkaConsumer{brokerAddress: brokerAddress, topic: topic}
}

// kafka consumer receive a kafka message and execute a callback with this message
func (k *kafkaConsumer) Consume(callback func([]byte)) {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{k.brokerAddress},
		Topic:   k.topic,
		GroupID: "mygroup",
	})
	defer reader.Close()

	fmt.Println("kafka start consuming on topic", k.topic)

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		callback(msg.Value)
	}
}
