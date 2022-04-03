package entity

import "github.com/google/uuid"

type Streamer struct {
	ID    uuid.UUID
	Name  string
	Image string
}

func NewStreamer(name string) Streamer {
	return Streamer{
		ID:   uuid.New(),
		Name: name,
	}
}
