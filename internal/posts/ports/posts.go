package ports

import (
	"github.com/erkylima/posts-service/internal/posts/domains"
	"github.com/erkylima/posts-service/internal/shared/entities"
)

type PostService interface {
	CreatePost(post *domains.Post) (string, error)
	ReadPost(slug string) (*domains.Post, error)
	UpdatePost(slug string, post *domains.Post) (*domains.Post, error)
	DeletePost(slug string) error
	ListPosts(filter []entities.Filter) ([]domains.Post, error)
}
