package entity

import "github.com/google/uuid"

type Streamer struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}

func NewStreamer(name string) *Streamer {
	return &Streamer{
		ID:   uuid.New(),
		Name: name,
	}
}
