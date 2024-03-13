package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type PostService struct {
	repo ports.PostRepository
}

func NewPostService(repo ports.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post domains.Post) (*domains.Post, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) ReadPost(slug string) (*domains.Post, error) {
	return s.repo.ReadPost(slug)
}

func (s *PostService) UpdatePost(post domains.Post) (*domains.Post, error) {
	return s.repo.UpdatePost(post)
}

func (s *PostService) DeletePost(slug string) error {
	return s.repo.DeletePost(slug)
}

func (s *PostService) ListPosts() ([]domains.Post, error) {
	return s.repo.ListPosts()
}

func (s *PostService) ListPostsByTag(tag string) ([]domains.Post, error) {
	return s.repo.ListPostsByTag(tag)
}

func (s *PostService) ListPostsByCategory(category string) ([]domains.Post, error) {
	return s.repo.ListPostsByCategory(category)
}

func (s *PostService) ListPostsByAuthor(author string) ([]domains.Post, error) {
	return s.repo.ListPostsByAuthor(author)
}

func (s *PostService) ListPostsByDate(date string) ([]domains.Post, error) {
	return s.repo.ListPostsByDate(date)
}
