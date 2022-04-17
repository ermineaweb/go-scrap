package main

import (
	"go-twitch/infra/kafka"
)

func main() {
	startStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_start")
	// stopStreamNotifProducer := kafka.NewKafkaProducer("kafka:9092", "stream_stop")
	for {
		startStreamNotifProducer.Produce([]byte("zerator"))
		startStreamNotifProducer.Produce([]byte("otplol_"))
		startStreamNotifProducer.Produce([]byte("mistermv"))
		startStreamNotifProducer.Produce([]byte("chap_gg"))
		startStreamNotifProducer.Produce([]byte("asmongold"))
		// time.Sleep(20 * time.Second)
		// stopStreamNotifProducer.Produce([]byte("zerator"))
		// time.Sleep(5 * time.Second)
	}
}
