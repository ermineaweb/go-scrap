package streamer

import (
	"go-scrap/domain/entity"

	"github.com/google/uuid"
)

type PostresqlRepository struct{}

func NewPostresqlRepository() *PostresqlRepository {
	return &PostresqlRepository{}
}

func (r *PostresqlRepository) Save(streamer entity.Streamer) {
	// todo
}

func (r *PostresqlRepository) GetById(ID uuid.UUID) (entity.Streamer, error) {
	// todo
	return entity.Streamer{}, nil
}

func (r *PostresqlRepository) GetAll() []entity.Streamer {
	// todo
	return []entity.Streamer{}
}

func (r *PostresqlRepository) Delete(ID uuid.UUID) error {
	// todo
	return nil
}
