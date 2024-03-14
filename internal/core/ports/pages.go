package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type PageService interface {
	CreatePage(page *domains.Page) (*domains.Page, error)
	ReadPage(slug string) (*domains.Page, error)
	UpdatePage(page domains.Page) (*domains.Page, error)
	DeletePage(slug string) error
	ListPages() ([]domains.Page, error)
	ListPagesByCategory(category string) ([]domains.Page, error)
	ListPagesByTag(tag string) ([]domains.Page, error)
	ListPagesByAuthor(author string) ([]domains.Page, error)
}
