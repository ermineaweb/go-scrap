package inmemory

import (
	"errors"
	"go-twitch/domain/entity"

	"github.com/google/uuid"
)

// var r.database = []entity.Streamer{
// 	{ID: uuid.New(), Name: "Harry", Image: ""},
// 	{ID: uuid.New(), Name: "Ron", Image: ""},
// }

var (
	ErrStreamerNotFound = errors.New("streamer not found")
)

type StreamerRepository struct {
	database []entity.Streamer
}

func NewStreamerRepository() *StreamerRepository {
	return &StreamerRepository{}
}

func (r *StreamerRepository) Save(streamer entity.Streamer) {
	r.database = append(r.database, streamer)
}

func (r *StreamerRepository) GetById(ID uuid.UUID) (entity.Streamer, error) {
	for _, streamer := range r.database {
		if streamer.ID == ID {
			return streamer, nil
		}
	}
	return entity.Streamer{}, ErrStreamNotFound
}

func (r *StreamerRepository) GetAll() []entity.Streamer {
	return r.database
}

func (r *StreamerRepository) Delete(ID uuid.UUID) error {
	// TODO
	return nil
}
