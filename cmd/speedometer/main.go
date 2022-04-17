package main

import (
	"fmt"
	"go-twitch/domain/services"
	"go-twitch/infra/kafka"
	"sort"
	"time"

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
	go speedometerService.Run()

	refreshDisplay := time.NewTicker(1 * time.Second)
	defer refreshDisplay.Stop()

	for range refreshDisplay.C {
		fmt.Print("\033[H\033[2J")
		sort.Sort(speedometerService.Chats)
		for _, chat := range speedometerService.Chats {
			fmt.Printf("%s\ntotal:\t%d\nspeed:\t%d msg/min\n\n",
				chat.Streamer.Name,
				chat.NbMessageOverStart,
				chat.MsgPerMin,
			)
		}
	}
}
