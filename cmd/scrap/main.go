package main

import (
	"fmt"
	"go-twitch/domain/entity"
	"go-twitch/domain/services"
	"go-twitch/infra/inmemory"
)

func main() {
	ron := entity.NewStreamer("ron")
	s := entity.NewStream(ron)
	fmt.Printf("%+v\n", s)
	streamRepo := inmemory.NewStreamRepository()
	streamerRepo := inmemory.NewStreamerRepository()
	scrapService := services.NewScrapingService(streamRepo, streamerRepo)
	scrapService.Scrap(ron)
}
