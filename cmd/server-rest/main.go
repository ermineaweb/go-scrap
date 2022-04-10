package main

import (
	"fmt"
	"go-twitch/domain/services"
	"go-twitch/infra/kafka"
	"go-twitch/infra/rest"
	"log"
	"net/http"

	"github.com/gempir/go-twitch-irc/v3"
)

var MEASURE_INTERVAL = 5

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
	go speedometerService.Run()

	server := http.NewServeMux()
	server.HandleFunc("/speed", rest.SpeedometerHandler(&speedometerService.Chats))

	fmt.Println("Serve on 3000...")
	log.Fatal(http.ListenAndServe(":3000", server))
}
