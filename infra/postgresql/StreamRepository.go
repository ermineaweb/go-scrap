package postgresql

import (
	"go-twitch/domain/entity"

	"github.com/google/uuid"
)

type StreamRepository struct{}

func NewStreamRepository() *StreamRepository {
	return &StreamRepository{}
}

func (r *StreamRepository) Save(stream entity.Stream) {
	// todo
}

func (r *StreamRepository) GetById(ID uuid.UUID) (entity.Stream, error) {
	// todo
	return entity.Stream{}, nil
}

func (r *StreamRepository) GetAll() []entity.Stream {
	// todo
	return []entity.Stream{}
}

func (r *StreamRepository) Delete(ID uuid.UUID) error {
	// todo
	return nil
}
