package ports

import (
	"github.com/erkylima/posts-service/internal/pages/domains"
	"github.com/erkylima/posts-service/internal/shared/entities"
)

type PageService interface {
	CreatePage(page *domains.Page) (string, error)
	ReadPage(slug string) (*domains.Page, error)
	UpdatePage(slug string, page *domains.Page) (*domains.Page, error)
	DeletePage(slug string) error
	ListPages(filter []entities.Filter) ([]domains.Page, error)
}
