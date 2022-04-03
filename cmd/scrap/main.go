package main

import (
	"fmt"
	"go-scrap/domain/entity"
	"go-scrap/domain/services"
	"go-scrap/infra/stream"
	"go-scrap/infra/streamer"
)

func main() {
	ron := entity.NewStreamer("ron")
	s := entity.NewStream(ron)
	fmt.Printf("%+v\n", s)
	streamRepo := stream.NewInMemoryRepository()
	streamerRepo := streamer.NewInMemoryRepository()
	scrapService := services.NewScrapingService(streamRepo, streamerRepo)
	scrapService.Scrap(ron)
}
