package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type PageService struct {
	repo ports.PageRepository
}

func NewPageService(repo ports.PageRepository) *PageService {
	return &PageService{repo: repo}
}

func (ps *PageService) CreatePage(page domains.Page) (domains.Page, error) {
	return ps.repo.CreatePage(page)
}

func (ps *PageService) ReadPage(slug string) (domains.Page, error) {
	return ps.repo.ReadPage(slug)
}

func (ps *PageService) UpdatePage(page domains.Page) (domains.Page, error) {
	return ps.repo.UpdatePage(page)
}

func (ps *PageService) DeletePage(slug string) error {
	return ps.repo.DeletePage(slug)
}

func (ps *PageService) ListPages() ([]domains.Page, error) {
	return ps.repo.ListPages()
}

func (ps *PageService) ListPagesByCategory(category string) ([]domains.Page, error) {
	return ps.repo.ListPagesByCategory(category)
}

func (ps *PageService) ListPagesByTag(tag string) ([]domains.Page, error) {
	return ps.repo.ListPagesByTag(tag)
}

func (ps *PageService) ListPagesByAuthor(author string) ([]domains.Page, error) {
	return ps.repo.ListPagesByAuthor(author)
}
