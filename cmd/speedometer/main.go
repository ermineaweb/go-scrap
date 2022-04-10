package main

import (
	"go-twitch/domain/services"
	"go-twitch/infra/kafka"

	"github.com/gempir/go-twitch-irc/v3"
)

const MEASURE_INTERVAL = 10

func main() {
	startKafkaConsumer := kafka.NewKafkaConsumer("kafka:9092", "stream_start")
	stopKafkaConsumer := kafka.NewKafkaConsumer("kafka:9092", "stream_stop")
	twitchClient := twitch.NewAnonymousClient()
	speedometerService := services.NewSpeedometerService(
		twitchClient,
		startKafkaConsumer,
		stopKafkaConsumer,
		MEASURE_INTERVAL,
	)
	speedometerService.Run()
}
