package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type PostService interface {
	CreatePost(post *domains.Post) (*domains.Post, error)
	ReadPost(slug string) (*domains.Post, error)
	UpdatePost(slug string, post *domains.Post) (*domains.Post, error)
	DeletePost(slug string) error
	ListPosts(filter []domains.Filter) ([]domains.Post, error)
}
