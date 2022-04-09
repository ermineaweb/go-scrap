package main

import (
	"go-twitch/infra/kafka"
	"time"
)

func main() {
	startStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_start")
	stopStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_stop")
	for {
		startStreamNotifProducer.Produce([]byte("zerator"))
		time.Sleep(20 * time.Second)
		stopStreamNotifProducer.Produce([]byte("zerator"))
		time.Sleep(10 * time.Second)
	}
}
