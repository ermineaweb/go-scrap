package streamer

import (
	"errors"
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

var inMemoryDb = []entity.Streamer{
	{ID: uuid.New(), Name: "Harry", Image: ""},
	{ID: uuid.New(), Name: "Ron", Image: ""},
}

var (
	ErrStreamNotFound = errors.New("streamer not found")
)

type InMemoryRepository struct{}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func (r *InMemoryRepository) Save(streamer entity.Streamer) {
	inMemoryDb = append(inMemoryDb, streamer)
}

func (r *InMemoryRepository) GetById(ID uuid.UUID) (entity.Streamer, error) {
	for _, streamer := range inMemoryDb {
		if streamer.ID == ID {
			return streamer, nil
		}
	}
	return entity.Streamer{}, ErrStreamNotFound
}

func (r *InMemoryRepository) GetAll() []entity.Streamer {
	return inMemoryDb
}

func (r *InMemoryRepository) Delete(ID uuid.UUID) error {
	// TODO
	return nil
}
