package entity

import (
	"github.com/google/uuid"
)

type Stream struct {
	ID           uuid.UUID
	Streamer     *Streamer
	ViewersCount int
}

func NewStream(streamer *Streamer) *Stream {
	return &Stream{
		ID:       uuid.New(),
		Streamer: streamer,
	}
}
