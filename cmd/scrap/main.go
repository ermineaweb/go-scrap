package main

import (
	"fmt"
	"go-scrap/domain/entity"
	"go-scrap/domain/services"
	"go-scrap/infra/inmemory"
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
