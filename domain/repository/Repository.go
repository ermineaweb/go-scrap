package repository

import (
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

type Storable interface {
	entity.Stream | entity.Streamer
}

type Repository[T Storable] interface {
	Save(T)
	GetById(ID uuid.UUID) (T, error)
	GetAll() []T
	Delete(ID uuid.UUID) error
}

type Stream interface {
	Repository[entity.Stream]
}

type Streamer interface {
	Repository[entity.Streamer]
}
