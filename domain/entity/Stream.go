package entity

import (
	"github.com/google/uuid"
)

type Stream struct {
	ID           uuid.UUID `json:"id"`
	Streamer     *Streamer `json:"streamer"`
	Chat         *Chat     `json:"chat"`
	ViewersCount int       `json:"viewers_count"`
}

func NewStream(streamer *Streamer) *Stream {
	return &Stream{
		ID:       uuid.New(),
		Streamer: streamer,
		Chat:     NewChat(streamer),
	}
}
