package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type CategoryService interface {
	CreateCategory(category domains.Category) (*domains.Category, error)
	ReadCategory(slug string) (*domains.Category, error)
	UpdateCategory(category domains.Category) (*domains.Category, error)
	DeleteCategory(slug string) error
	ListCategories() ([]domains.Category, error)
}
