package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type PageService interface {
	CreatePage(page *domains.Page) (*domains.Page, error)
	ReadPage(slug string) (*domains.Page, error)
	UpdatePage(slug string, page *domains.Page) (*domains.Page, error)
	DeletePage(slug string) error
	ListPages(filter []domains.Filter) ([]domains.Page, error)
}
