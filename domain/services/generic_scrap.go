package services

// import (
// 	"fmt"
// 	"go-scrap/entity"

// 	"github.com/google/uuid"
// )

// type Scrapable interface {
// 	entity.Stream | entity.Streamer
// }

// type Repository[T Scrapable] interface {
// 	Save(T) uuid.UUID
// 	GetById(ID uuid.UUID) (*T, error)
// 	GetAll() []*T
// 	Delete(ID uuid.UUID) error
// }

// type ScrapService[T Scrapable] struct {
// 	repository Repository[T]
// }

// func NewScrapService[T Scrapable](repository Repository[T]) *ScrapService[T] {
// 	fmt.Println("create GENERIC scrap service")
// 	return &ScrapService[T]{repository: repository}
// }

// func (s *ScrapService[Scrapable]) Scrap(T Scrapable) {
// 	fmt.Println("we scrap")
// 	s.repository.Save(T)
// }

// func (s *ScrapService[Scrapable]) GetAll() []*Scrapable {
// 	return s.repository.GetAll()
// }
