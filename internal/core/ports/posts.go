package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type PostService interface {
	CreatePost(post domains.Post) (domains.Post, error)
	ReadPost(slug string) (domains.Post, error)
	UpdatePost(post domains.Post) (domains.Post, error)
	DeletePost(slug string) error
	ListPosts() ([]domains.Post, error)
	ListPostsByTag(tag string) ([]domains.Post, error)
	ListPostsByCategory(category string) ([]domains.Post, error)
	ListPostsByAuthor(author string) ([]domains.Post, error)
	ListPostsByDate(date string) ([]domains.Post, error)
}
