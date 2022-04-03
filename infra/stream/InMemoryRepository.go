package stream

import (
	"errors"
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

var inMemoryDb = []entity.Stream{
	entity.NewStream(entity.NewStreamer("harry")),
	entity.NewStream(entity.NewStreamer("hermione")),
}

var (
	ErrStreamNotFound = errors.New("stream not found")
)

type InMemoryRepository struct{}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (r *InMemoryRepository) Save(stream entity.Stream) {
	inMemoryDb = append(inMemoryDb, stream)
}

func (r *InMemoryRepository) GetById(ID uuid.UUID) (entity.Stream, error) {
	for _, stream := range inMemoryDb {
		if stream.ID == ID {
			return stream, nil
		}
	}
	return entity.Stream{}, ErrStreamNotFound
}

func (r *InMemoryRepository) GetAll() []entity.Stream {
	return inMemoryDb
}

func (r *InMemoryRepository) Delete(ID uuid.UUID) error {
	// TODO
	return nil
}
