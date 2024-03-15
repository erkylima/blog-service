package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type Repository[T any] interface {
	Create(entity *T) (string, error)
	Read(slug string, entity *T) error
	Update(slug string, entity *T) error
	Delete(slug string) error
	List(filter []domains.Filter) ([]T, error)
}
