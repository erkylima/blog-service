package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type CategoryService struct {
	repo ports.CategoryRepository
}

func NewCategoryRepository(repo ports.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cr *CategoryService) CreateCategory(category domains.Category) (*domains.Category, error) {
	return cr.repo.CreateCategory(category)
}

func (cr *CategoryService) ReadCategory(slug string) (*domains.Category, error) {
	return cr.repo.ReadCategory(slug)
}

func (cr *CategoryService) UpdateCategory(category domains.Category) (*domains.Category, error) {
	return cr.repo.UpdateCategory(category)
}

func (cr *CategoryService) DeleteCategory(slug string) error {
	return cr.repo.DeleteCategory(slug)
}

func (cr *CategoryService) ListCategories() ([]domains.Category, error) {
	return cr.repo.ListCategories()
}
