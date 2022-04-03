package stream

import (
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

type PostresqlRepository struct{}

func NewPostresqlRepository() *PostresqlRepository {
	return &PostresqlRepository{}
}

func (r *PostresqlRepository) Save(stream entity.Stream) {
	// todo
}

func (r *PostresqlRepository) GetById(ID uuid.UUID) (entity.Stream, error) {
	// todo
	return entity.Stream{}, nil
}

func (r *PostresqlRepository) GetAll() []entity.Stream {
	// todo
	return []entity.Stream{}
}

func (r *PostresqlRepository) Delete(ID uuid.UUID) error {
	// todo
	return nil
}
