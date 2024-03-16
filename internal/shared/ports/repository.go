package ports

import (
	"github.com/erkylima/posts-service/internal/shared/entities"
)

type Repository[T any] interface {
	Create(entity *T) (string, error)
	Read(slug string, entity *T) error
	Update(slug string, entity *T) error
	Delete(slug string) error
	List(filter []entities.Filter) ([]T, error)
}
