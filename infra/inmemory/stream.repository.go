package inmemory

import (
	"errors"
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

// var r.database = []entity.Stream{
// 	entity.NewStream(entity.NewStreamer("harry")),
// 	entity.NewStream(entity.NewStreamer("hermione")),
// }

var (
	ErrStreamNotFound = errors.New("stream not found")
)

type StreamRepository struct {
	database []entity.Stream
}

func NewStreamRepository() *StreamRepository {
	return &StreamRepository{}
}

func (r *StreamRepository) Save(stream entity.Stream) {
	r.database = append(r.database, stream)
}

func (r *StreamRepository) GetById(ID uuid.UUID) (entity.Stream, error) {
	for _, stream := range r.database {
		if stream.ID == ID {
			return stream, nil
		}
	}
	return entity.Stream{}, ErrStreamNotFound
}

func (r *StreamRepository) GetAll() []entity.Stream {
	return r.database
}

func (r *StreamRepository) Delete(ID uuid.UUID) error {
	// TODO
	return nil
}
