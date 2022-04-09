package postgresql

import (
	"go-twitch/domain/entity"

	"github.com/google/uuid"
)

type StreamerRepository struct{}

func NewStreamerRepository() *StreamerRepository {
	return &StreamerRepository{}
}

func (r *StreamerRepository) Save(streamer entity.Streamer) {
	// todo
}

func (r *StreamerRepository) GetById(ID uuid.UUID) (entity.Streamer, error) {
	// todo
	return entity.Streamer{}, nil
}

func (r *StreamerRepository) GetAll() []entity.Streamer {
	// todo
	return []entity.Streamer{}
}

func (r *StreamerRepository) Delete(ID uuid.UUID) error {
	// todo
	return nil
}
