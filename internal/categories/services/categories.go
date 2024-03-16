package services

import (
	"github.com/erkylima/posts-service/internal/categories/domains"
	"github.com/erkylima/posts-service/internal/shared/ports"
)

type CategoryService struct {
	repo ports.Repository[domains.Category]
}

func NewCategoryRepository(repo ports.Repository[domains.Category]) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cr *CategoryService) CreateCategory(category domains.Category) (*domains.Category, error) {
	// cr.repo.Create(category)
	return nil, nil
}

func (cr *CategoryService) ReadCategory(slug string) (*domains.Category, error) {
	// return cr.repo.ReadCategory(slug)
	return nil, nil
}

func (cr *CategoryService) UpdateCategory(category domains.Category) (*domains.Category, error) {
	// return cr.repo.UpdateCategory(category)
	return nil, nil
}

func (cr *CategoryService) DeleteCategory(slug string) error {
	// return cr.repo.DeleteCategory(slug)
	return nil
}

func (cr *CategoryService) ListCategories() ([]domains.Category, error) {
	// return cr.repo.ListCategories()
	return nil, nil
}
