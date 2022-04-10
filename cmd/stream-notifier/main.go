package main

import (
	"go-twitch/infra/kafka"
	"time"
)

func main() {
	startStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_start")
	// stopStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_stop")
	for {
		startStreamNotifProducer.Produce([]byte("zerator"))
		startStreamNotifProducer.Produce([]byte("otplol_"))
		startStreamNotifProducer.Produce([]byte("thebausffs"))
		startStreamNotifProducer.Produce([]byte("tubbo"))
		// time.Sleep(20 * time.Second)
		// stopStreamNotifProducer.Produce([]byte("zerator"))
		time.Sleep(10 * time.Second)
	}
}
