package services

import (
	"go-twitch/domain/entity"
	"go-twitch/domain/repository"
)

type StreamerService struct {
	repository repository.Streamer
}

func NewStreamerService(repository repository.Streamer) *StreamerService {
	return &StreamerService{repository: repository}
}

func (s StreamerService) GetAll() []entity.Streamer {
	return s.repository.GetAll()
}
