package repository

import "github.com/erkylima/posts-service/internal/core/domains"

type pageRepository struct {
}

func NewPageRepository() *pageRepository {
	return &pageRepository{}
}

func (r *pageRepository) Create(page *domains.Page) (*domains.Page, error) {

	return page, nil
}

func (r *pageRepository) Read(slug string, page *domains.Page) (*[]byte, error) {

	return nil, nil
}

func (r *pageRepository) Update(entity interface{}) (*[]byte, error) {
	return nil, nil
}
func (r *pageRepository) Delete(slug string) error {
	return nil
}
func (r *pageRepository) List(entity interface{}) ([]byte, error) {
	return nil, nil
}

func (r *pageRepository) ListByTag(tag string, entity interface{}) ([]byte, error) {
	return nil, nil
}
func (r *pageRepository) ListByCategory(category string, entity interface{}) ([]byte, error) {
	return nil, nil
}

func (r *pageRepository) ListByAuthor(author string, entity interface{}) ([]byte, error) {
	return nil, nil
}

func (r *pageRepository) ListByDate(date string, entity interface{}) ([]byte, error) {
	return nil, nil
}
