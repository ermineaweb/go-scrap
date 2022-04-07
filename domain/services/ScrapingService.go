package services

import (
	"fmt"
	"go-twitch/domain/entity"
	"go-twitch/domain/repository"
)

type ScrapingService struct {
	streamRepo   repository.Stream
	streamerRepo repository.Streamer
}

func NewScrapingService(streamRepo repository.Stream, streamerRepo repository.Streamer) ScrapingService {
	fmt.Println("create scrap service")
	return ScrapingService{streamRepo, streamerRepo}
}

func (s *ScrapingService) ScrapAll(stream entity.Stream) {
	streamers := NewStreamerService(s.streamerRepo).GetAll()
	for _, streamer := range streamers {
		s.Scrap(streamer)
	}
}

func (s *ScrapingService) Scrap(streamer entity.Streamer) {
	fmt.Println(streamer.Name)
}
