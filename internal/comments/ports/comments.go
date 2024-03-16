package ports

import "github.com/erkylima/posts-service/internal/comments/domains"

type CategoryService interface {
	CreateCategory(category domains.Comment) (*domains.Comment, error)
	ReadCategory(slug string) (*domains.Comment, error)
	UpdateCategory(category domains.Comment) (*domains.Comment, error)
	DeleteCategory(slug string) error
	ListCategories() ([]domains.Comment, error)
}
