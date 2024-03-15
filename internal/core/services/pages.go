package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type PageService struct {
	repo ports.Repository[domains.Page]
}

func NewPageService(repo ports.Repository[domains.Page]) *PageService {
	return &PageService{repo: repo}
}

func (ps *PageService) CreatePage(page *domains.Page) (string, error) {
	result, err := ps.repo.Create(page)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (ps *PageService) ReadPage(slug string) (*domains.Page, error) {
	page := &domains.Page{}
	err := ps.repo.Read(slug, page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (ps *PageService) UpdatePage(slug string, page *domains.Page) (*domains.Page, error) {
	err := ps.repo.Update(slug, page)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (ps *PageService) DeletePage(slug string) error {
	err := ps.repo.Delete(slug)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PageService) ListPages(filter []domains.Filter) ([]domains.Page, error) {

	pages, err := ps.repo.List(filter)
	if err != nil {
		return nil, err
	}

	return pages, nil
}
